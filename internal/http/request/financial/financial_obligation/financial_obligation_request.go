package financialobligationrequest

import (
	"finance/internal/apperrors"
	mxl "finance/internal/constants/max_len"
	financialmodel "finance/models/financial"
	"fmt"
)

type FinancialObligationRequest struct {
	CategoryId  int                                       `json:"category_id" validate:"required"`
	Description string                                    `json:"description" validate:"required"`
	Type        financialmodel.FinancialObligationsTpyes  `json:"type" validate:"required"`
	Status      financialmodel.FinancialObligationsStatus `json:"status" validate:"required"`
}

func validateFinancialObligationsTpyes(t financialmodel.FinancialObligationsTpyes) bool {
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

	if !validateFinancialObligationsTpyes(f.Type) {
		errors["type"] = append(errors["type"], "O tipo da obrigação financeira está inválido.")
	}

	return errors
}
