package financialobligationvalidator

import (
	"context"
	"errors"
	"finance/internal/apperrors"
	mxl "finance/internal/constants/max_len"
	financialobligationrequest "finance/internal/http/request/financial/financial_obligation"
	"finance/internal/logger"
	financialaccountrepository "finance/internal/modules/financial_accounts/repository"
	financialmodel "finance/models/financial"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type FinancialObligationRepository interface {
	ValidCategoryType(ctx context.Context, categoryId int, movementType financialmodel.FinancialObligationsTypes) (bool, error)
	FindById(ctx context.Context, financialObligationId int) (financialmodel.FinancialObligationModel, error)
	CalculateOutstandingBalance(ctx context.Context, financialObligationId int) (float64, error)
}

type FinancialObligationValidator struct {
	repo                       FinancialObligationRepository
	financialAccountRepository *financialaccountrepository.FinancialAccountRepository
}

func NewFinancialObligationValidatorValidator(repo FinancialObligationRepository, financialAccountRepository *financialaccountrepository.FinancialAccountRepository) *FinancialObligationValidator {
	return &FinancialObligationValidator{
		repo:                       repo,
		financialAccountRepository: financialAccountRepository,
	}
}

func parseDates(date string) (*time.Time, error) {
	parsedDate, err := time.Parse("2006-01-02", date)

	if err != nil {
		logger.General.Error.Println("Erro ao converter a data informada:", err)
		return nil, err

	}

	return &parsedDate, nil
}

func ValidateFinancialObligationsTypes(t financialmodel.FinancialObligationsTypes) bool {
	switch t {
	case financialmodel.PAYABLE,
		financialmodel.RECEIVABLE:

		return true

	default:
		return false
	}
}

func (f *FinancialObligationValidator) validateCategory(ctx context.Context, categoryId int, movementType financialmodel.FinancialObligationsTypes) error {
	isInvalidType, err := f.repo.ValidCategoryType(ctx, categoryId, movementType)

	if err != nil {
		logger.General.Error.Println("Erro ao válidar se o tipo da categoria é coerente com o tipo de movimento financeiro:", err)
		return fmt.Errorf("Erro ao válidar se o tipo da categoria é coerente com o tipo de movimento financeiro: %s", err)
	}

	if !isInvalidType {
		return fmt.Errorf("Tipo da categoria incoerente com o tipo da obrigação financeira")
	}

	return nil
}

func (f *FinancialObligationValidator) ValidatePayload(ctx context.Context, payload financialobligationrequest.FinancialObligationRequest) error {

	logger.General.Info.Println("---- Vai validar o payload do movimento financeiro via db ----")

	errors := apperrors.ValidationErrors{}

	err := f.validateCategory(ctx, payload.CategoryId, payload.Type)

	//category_id
	if err != nil {
		logger.General.Error.Println("Erro ao válidar se o tipo da categoria é coerente com o tipo de movimento financeiro:", err)
		errors["category_id"] = append(errors["category_id"], err.Error())
	}

	//description
	if payload.Description == "" {
		errors["description"] = append(errors["description"], "A descrição da obrigação financeira é obrigatório.")
	} else if len(payload.Description) > mxl.MAX_LEN_255 {
		errors["description"] = append(errors["description"], fmt.Sprintf("A descrição da obrigação financeira deve ter no máximo %d caracteres.", mxl.MAX_LEN_255))
	}

	//type
	if !ValidateFinancialObligationsTypes(payload.Type) {
		errors["type"] = append(errors["type"], "O tipo da obrigação financeira está inválido.")
	}

	//original_amount
	if payload.OriginalAmount <= 0 {
		errors["original_amount"] = append(errors["original_amount"], "O valor da obrigação financeira precisa ser maior que zero..")
	}

	//due_date
	if payload.DueDate == "" {
		errors["due_date"] = append(errors["due_date"], "A data de vencimento é obrigatória.")
	}

	if _, err := time.Parse("2006-01-02", payload.DueDate); err != nil {
		errors["due_date"] = append(errors["due_date"], "A data de vencimento deve estar no formato YYYY-MM-DD.")
	}

	//competence_date
	if payload.CompetenceDate != nil {
		if _, err := time.Parse("2006-01-02", *payload.CompetenceDate); err != nil {
			errors["competence_date"] = append(errors["competence_date"], "A data de competência deve estar no formato YYYY-MM-DD.")
		}
	}

	//notes
	if payload.Notes != nil {
		if len(*payload.Notes) > mxl.MAX_LEN_500 {
			errors["notes"] = append(errors["notes"], fmt.Sprintf("As notas da obrigação financeira deve ter no máximo %d caracteres.", mxl.MAX_LEN_500))
		}
	}

	logger.General.Info.Printf("---- Terminou de validar o payload da transação financeira, total de erros: %d ----", len(errors))

	if len(errors) > 0 {
		return apperrors.NewValidationError(errors)
	}

	return nil
}

func (f FinancialObligationValidator) ValidateUpdatePayload(ctx context.Context, payload financialobligationrequest.FinancialObligationRequest, financialObligationId int) error {
	errors := apperrors.ValidationErrors{}

	financialObligation, err := f.repo.FindById(ctx, financialObligationId)

	if err != nil {
		logger.General.Error.Println("Erro ao válidar se a obrigação financeira está parcialmente paga:", err)
		return err
	}

	if financialObligation.Status == financialmodel.CANCELED {
		errors["status"] = append(errors["status"], "Essa obrigação financeira está cancelada.")
		return apperrors.NewValidationError(errors)
	}

	if financialObligation.Status == financialmodel.SETTLED {
		errors["status"] = append(errors["status"], "Essa obrigação financeira está paga.")
		return apperrors.NewValidationError(errors)
	}

	if financialObligation.Status == financialmodel.PENDING {
		//category_id
		err := f.validateCategory(ctx, payload.CategoryId, payload.Type)

		if err != nil {
			logger.General.Error.Println("Erro ao válidar se o tipo da categoria é coerente com o tipo de movimento financeiro:", err)
			errors["category_id"] = append(errors["category_id"], err.Error())
		}

		//description
		if payload.Description == "" {
			errors["description"] = append(errors["description"], "A descrição da obrigação financeira é obrigatório.")
		} else if len(payload.Description) > mxl.MAX_LEN_255 {
			errors["description"] = append(errors["description"], fmt.Sprintf("A descrição da obrigação financeira deve ter no máximo %d caracteres.", mxl.MAX_LEN_255))
		}

		//type
		if !ValidateFinancialObligationsTypes(payload.Type) {
			errors["type"] = append(errors["type"], "O tipo da obrigação financeira está inválido.")
		}

		//original_amount
		if payload.OriginalAmount <= 0 {
			errors["original_amount"] = append(errors["original_amount"], "O valor da obrigação financeira precisa ser maior que zero..")
		}

		//due_date
		if payload.DueDate == "" {
			errors["due_date"] = append(errors["due_date"], "A data de vencimento é obrigatória.")
		}

		if _, err := time.Parse("2006-01-02", payload.DueDate); err != nil {
			errors["due_date"] = append(errors["due_date"], "A data de vencimento deve estar no formato YYYY-MM-DD.")
		}

		//competence_date
		if payload.CompetenceDate != nil {
			if _, err := time.Parse("2006-01-02", *payload.CompetenceDate); err != nil {
				errors["competence_date"] = append(errors["competence_date"], "A data de competência deve estar no formato YYYY-MM-DD.")
			}
		}

		//notes
		if payload.Notes != nil {
			if len(*payload.Notes) > mxl.MAX_LEN_500 {
				errors["notes"] = append(errors["notes"], fmt.Sprintf("As notas da obrigação financeira deve ter no máximo %d caracteres.", mxl.MAX_LEN_500))
			}
		}
	}

	if financialObligation.Status == financialmodel.PARTIALLY_SETTLED {

		//description
		if payload.Description == "" {
			errors["description"] = append(errors["description"], "A descrição da obrigação financeira é obrigatório.")
		} else if len(payload.Description) > mxl.MAX_LEN_255 {
			errors["description"] = append(errors["description"], fmt.Sprintf("A descrição da obrigação financeira deve ter no máximo %d caracteres.", mxl.MAX_LEN_255))
		}

		//notes
		if payload.Notes != nil {
			if len(*payload.Notes) > mxl.MAX_LEN_500 {
				errors["notes"] = append(errors["notes"], fmt.Sprintf("As notas da obrigação financeira deve ter no máximo %d caracteres.", mxl.MAX_LEN_500))
			}
		}
	}

	if len(errors) > 0 {
		return apperrors.NewValidationError(errors)
	}

	return nil
}

func (f FinancialObligationValidator) ValidateCancel(ctx context.Context, financialObligationId int) error {
	errors := apperrors.ValidationErrors{}

	financialObligation, err := f.repo.FindById(ctx, financialObligationId)

	if err != nil {
		logger.General.Error.Println("Erro ao válidar se a obrigação financeira existe:", err)
		return err
	}

	switch financialObligation.Status {
	case financialmodel.SETTLED:
		errors["status"] = append(errors["status"], "Essa obrigação financeira está paga.")
		return apperrors.NewValidationError(errors)

	case financialmodel.PARTIALLY_SETTLED:
		errors["status"] = append(errors["status"], "Essa obrigação financeira está parcialmente paga.")
		return apperrors.NewValidationError(errors)

	case financialmodel.CANCELED:
		errors["status"] = append(errors["status"], "Essa obrigação financeira já cancelada.")
		return apperrors.NewValidationError(errors)
	}

	if len(errors) > 0 {
		return apperrors.NewValidationError(errors)
	}

	return nil
}

func (f FinancialObligationValidator) ValidatePayObligationPayload(ctx context.Context, payload financialobligationrequest.PayFinancialObligationRequest) error {
	validationErrors := apperrors.ValidationErrors{}

	financialObligation, err := f.repo.FindById(ctx, payload.FinancialObligationId)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) || errors.Is(err, apperrors.ErrNotFound) {
			validationErrors["financial_obligation_id"] = append(validationErrors["financial_obligation_id"], "Obrigação financeira não localizada.")
			return apperrors.NewValidationError(validationErrors)
		}

		logger.General.Error.Println("Erro ao válidar se a obrigação financeira existe:", err)
		return err
	}

	if _, err = f.financialAccountRepository.FindById(ctx, payload.FinancialAccountId); err != nil {
		if errors.Is(err, pgx.ErrNoRows) || errors.Is(err, apperrors.ErrNotFound) {
			validationErrors["financial_account_id"] = append(validationErrors["financial_account_id"], "Conta financeira não localizada.")
			return apperrors.NewValidationError(validationErrors)
		}

		logger.General.Error.Println("Erro ao válidar se a obrigação financeira existe:", err)
		return err
	}

	if financialObligation.Type != financialmodel.PAYABLE {
		validationErrors["type"] = append(
			validationErrors["type"],
			"Somente obrigações a pagar podem ser pagas.",
		)
	}

	switch financialObligation.Status {
	case financialmodel.SETTLED:
		validationErrors["status"] = append(validationErrors["status"], "Essa obrigação financeira já está paga.")
		return apperrors.NewValidationError(validationErrors)

	case financialmodel.CANCELED:
		validationErrors["status"] = append(validationErrors["status"], "Essa obrigação financeira está cancelada.")
		return apperrors.NewValidationError(validationErrors)
	}

	if payload.AmountPaid <= 0 {
		validationErrors["amount_paid"] = append(validationErrors["amount_paid"], "O valor pago precisa ser maior que zero.")
	}

	outstandingBalance, err := f.repo.CalculateOutstandingBalance(ctx, payload.FinancialObligationId)

	if err != nil {
		return err
	}

	if payload.AmountPaid > outstandingBalance {
		validationErrors["amount_paid"] = append(validationErrors["amount_paid"], "O valor pago não pode ser maior que o saldo pendente da obrigação.")
	}

	if payload.PaymentDate == "" {
		validationErrors["payment_date"] = append(validationErrors["payment_date"], "A data de pagamento é obrigatória.")
	} else if _, err := time.Parse("2006-01-02", payload.PaymentDate); err != nil {
		validationErrors["payment_date"] = append(validationErrors["payment_date"], "A data de pagamento deve estar no formato YYYY-MM-DD.")
	}

	if len(validationErrors) > 0 {
		logger.General.Error.Println(validationErrors)

		return apperrors.NewValidationError(validationErrors)
	}

	return nil
}
