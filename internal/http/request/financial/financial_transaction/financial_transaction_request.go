package financialtransactionrequest

import (
	"finance/internal/apperrors"
	mxl "finance/internal/constants/max_len"
	"finance/internal/logger"
	financialmodel "finance/models/financial"
	"fmt"
	"time"
)

type FinancialTransactionRequest struct {
	FinancialAccountId    int                                               `json:"financial_account_id" validate:"required"`
	CategoryId            int                                               `json:"category_id" validate:"required"`
	ReversedTransactionId *int                                              `json:"reversed_transaction_id"`
	OriginType            *string                                           `json:"origin_type"`
	OriginId              *int                                              `json:"origin_id"`
	Description           string                                            `json:"description" validate:"required"`
	Amount                float64                                           `json:"amount" validate:"required"`
	MovementDate          string                                            `json:"movement_date" validate:"required"`
	ReferenceDate         string                                            `json:"reference_date" validate:"required"`
	MovementType          financialmodel.FinancialTransactionsMovementType  `json:"movement_type" validate:"required"`
	OperationType         financialmodel.FinancialTransactionsOperationType `json:"operation_type" validate:"required"`
}

func validateFinancialTransactionsMovementType(t financialmodel.FinancialTransactionsMovementType) bool {
	switch t {
	case financialmodel.ENTRY,
		financialmodel.EXIT:

		return true

	default:
		return false
	}
}

func validateFinancialTransactionsOperationType(t financialmodel.FinancialTransactionsOperationType) bool {
	switch t {
	case financialmodel.ORIGINAL,
		financialmodel.ADJUSTMENT,
		financialmodel.REVERSAL:

		return true

	default:
		return false
	}
}

func (f FinancialTransactionRequest) ValidatePayload() apperrors.ValidationErrors {
	logger.General.Info.Println("---- Vai validar o payload da movimentação financeira via request ----")

	errors := apperrors.ValidationErrors{}

	// FinancialAccountId
	if f.FinancialAccountId < 0 {
		errors["financial_account_id"] = append(errors["financial_account_id"], "O ID de referência da conta financeira não pode ser menor que zero.")
	}

	//CategoryId category_id
	if f.CategoryId < 0 {
		errors["category_id"] = append(errors["parent_id"], "O ID de referência da categoria não pode ser menor que zero.")
	}

	//Description description
	if f.Description == "" {
		errors["description"] = append(errors["description"], "A descrição da movimentação financeira é obrigatório.")
	} else if len(f.Description) > mxl.MAX_LEN_255 {
		errors["description"] = append(errors["description"], fmt.Sprintf("A descrição da movimentação financeira deve ter no máximo %d caracteres.", mxl.MAX_LEN_255))
	}

	//Amount amount
	if f.Amount <= 0 {
		errors["amount"] = append(errors["amount"], "O valor da movimentação financeira não pode ser menor que zero.")
	}

	//MovementDate movement_date
	if _, err := time.Parse("2006-01-02", f.MovementDate); err != nil {
		errors["movement_date"] = append(errors["movement_date"], "A data de movimentação deve estar no formato YYYY-MM-DD.")
	}

	//ReferenceDate reference_date
	if _, err := time.Parse("2006-01-02", f.ReferenceDate); err != nil {
		errors["reference_date"] = append(errors["reference_date"], "A data de referência deve estar no formato YYYY-MM-DD.")
	}

	//IdempotencyKey idempotency_key

	//MovementType movement_type
	if !validateFinancialTransactionsMovementType(f.MovementType) {
		errors["movement_type"] = append(errors["movement_type"], "O tipo da movimentação financeira está inválido.")
	}

	//OperationType operation_type
	if !validateFinancialTransactionsOperationType(f.OperationType) {
		errors["operation_type"] = append(errors["operation_type"], "O tipo da operação financeira está inválido.")
	}

	logger.General.Info.Printf("---- Terminou de validar o payload da movimentação financeira, total de erros: %d ----", len(errors))
	return errors
}
