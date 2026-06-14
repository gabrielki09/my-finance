package categoryrequest

import (
	"finance/internal/apperrors"
	mxl "finance/internal/constants/max_len"
	"finance/internal/logger"
	categorymodel "finance/models/category"
	"fmt"
	"strings"
)

type CategoryRequest struct {
	Id       int
	ParentId *int                        `json:"parent_id" validate:"sometimes,numeric"`
	Name     string                      `json:"name" validate:"required"`
	Type     categorymodel.CategoryTpyes `json:"type" validate:"required"`
}

func validateType(t categorymodel.CategoryTpyes) bool {
	switch t {
	case categorymodel.INCOME,
		categorymodel.EXPENSE,
		categorymodel.BOTH:

		return true

	default:
		return false
	}
}

func (c CategoryRequest) ValidatePayload() apperrors.ValidationErrors {
	logger.General.Info.Println("---- Vai validar o payload da categoria via request ----")

	errors := apperrors.ValidationErrors{}

	name := strings.TrimSpace(c.Name)
	categoryType := categorymodel.CategoryTpyes(c.Type)

	if name == "" {
		errors["name"] = append(errors["name"], "O nome da categoria é obrigatório.")
	} else if len(name) > mxl.MAX_LEN_100 {
		errors["name"] = append(errors["name"], fmt.Sprintf("O nome da categoria deve ter no máximo %d caracteres.", mxl.MAX_LEN_100))
	}

	if categoryType == "" {
		errors["type"] = append(errors["type"], "O tipo da categoria é obrigatório.")
	} else if !validateType(categoryType) {
		errors["type"] = append(errors["type"], "O tipo da categoria é inválido.")
	}

	if c.ParentId != nil {
		if *c.ParentId < 0 {
			errors["parent_id"] = append(errors["parent_id"], "O ID de referência não pode ser menor que zero.")
		}

		if *c.ParentId == c.Id {
			errors["parent_id"] = append(errors["parent_id"], "O ID de referência não pode ser o mesmo ID do registro.")
		}
	}

	logger.General.Info.Printf("---- Terminou de validar o payload da categoria, total de erros: %d ----", len(errors))
	return errors
}
