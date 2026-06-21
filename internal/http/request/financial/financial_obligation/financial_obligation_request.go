package financialobligationrequest

import (
	"finance/internal/apperrors"
	mxl "finance/internal/constants/max_len"
	financialmodel "finance/models/financial"
	"fmt"
	"time"
)

type FinancialObligationRequest struct {
	CategoryId     int                                      `json:"category_id" validate:"required"`
	Description    string                                   `json:"description" validate:"required"`
	Type           financialmodel.FinancialObligationsTypes `json:"type" validate:"required"`
	OriginalAmount float64                                  `json:"original_amount" validate:"required"`
	DueDate        string                                   `json:"due_date" validate:"required"`
	CompetenceDate *string                                  `json:"competence_date" validate:"required"`
	Notes          *string                                  `json:"notes" validate:"required"`
}

func validateFinancialObligationsTypes(t financialmodel.FinancialObligationsTypes) bool {
	switch t {
	case financialmodel.PAYABLE,
		financialmodel.RECEIVABLE:

		return true

	default:
		return false
	}
}

func (f FinancialObligationRequest) ValidatePayload() apperrors.ValidationErrors {
	errors := apperrors.ValidationErrors{}

	if f.CategoryId < 0 {
		errors["category_id"] = append(errors["parent_id"], "O ID de referência da categoria não pode ser menor que zero.")
	}

	if f.Description == "" {
		errors["description"] = append(errors["description"], "A descrição da obrigação financeira é obrigatório.")
	} else if len(f.Description) > mxl.MAX_LEN_255 {
		errors["description"] = append(errors["description"], fmt.Sprintf("A descrição da obrigação financeira deve ter no máximo %d caracteres.", mxl.MAX_LEN_255))
	}

	if !validateFinancialObligationsTypes(f.Type) {
		errors["type"] = append(errors["type"], "O tipo da obrigação financeira está inválido.")
	}

	if f.OriginalAmount <= 0 {
		errors["original_amount"] = append(errors["original_amount"], "O valor da obrigação financeira precisa ser maior que zero..")
	}

	if f.CompetenceDate != nil {
		if _, err := time.Parse("2006-01-02", *f.CompetenceDate); err != nil {
			errors["competence_date"] = append(errors["competence_date"], "A data de competência deve estar no formato YYYY-MM-DD.")
		}
	}

	if _, err := time.Parse("2006-01-02", f.DueDate); err != nil {
		errors["due_date"] = append(errors["due_date"], "A data de vencimento deve estar no formato YYYY-MM-DD.")
	}

	return errors
}
