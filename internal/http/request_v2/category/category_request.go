package category

import (
	"context"
	"finance/internal/apperrors"
	mxl "finance/internal/constants/max_len"
	"finance/internal/logger"
	categorymodel "finance/models/category"
	"fmt"
	"strings"
)

type CategoryRepository interface {
	VerifyParentId(ctx context.Context, parentId int) (bool, error)
	VerifyExistsCategoryName(ctx context.Context, categoryName string) (*categorymodel.CategoryModel, error)
}

type CategoryRequest struct {
	Id       int
	ParentId *int                        `json:"parent_id" validate:"sometimes,numeric"`
	Name     string                      `json:"name" validate:"required"`
	Type     categorymodel.CategoryTpyes `json:"type" validate:"required"`
	repo     CategoryRepository
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

func (c CategoryRequest) ValidatePayload(ctx context.Context) apperrors.ValidationErrors {
	logger.Info("---- Vai validar o payload inicial da categoria ----")

	errors := apperrors.ValidationErrors{}

	name := strings.TrimSpace(c.Name)
	categoryType := categorymodel.CategoryTpyes(c.Type)

	if name == "" {
		errors["name"] = append(errors["name"], "O nome da categoria é obrigatório.")
	} else if len(name) > mxl.MAX_LEN_100 {
		errors["name"] = append(errors["name"], fmt.Sprintf("O nome da categoria deve ter no máximo %d caracteres.", mxl.MAX_LEN_100))
	}

	categoryByName, err := c.repo.VerifyExistsCategoryName(ctx, c.Name)

	if err != nil {
		logger.Error("Erro ao conferir se a categoria já existe pelo nome: ", err)

		errors["name"] = append(errors["name"], "Erro ao conferir se a categoria já existe pelo nome")
		return errors
	}

	if categoryByName != nil {
		errors["name"] = append(errors["name"], fmt.Sprintf("A categoria %s já existe, ID %d.", c.Name, categoryByName.Id))
		return errors
	}

	if categoryType == "" {
		errors["type"] = append(errors["type"], "O tipo da categoria é obrigatório.")
	} else if !validateType(categoryType) {
		errors["type"] = append(errors["type"], "O tipo da categoria é inválido.")
	}

	if c.ParentId != nil {
		if *c.ParentId == c.Id {
			errors["parent_id"] = append(errors["parent_id"], "O ID da categoria pai não pode ser o mesmo ID do registro.")
		}
	}

	if c.ParentId != nil {
		exists, err := c.repo.VerifyParentId(ctx, *c.ParentId)

		if err != nil {
			errors["parent_id"] = append(errors["parent_id"], "O ID da categoria pai não pode ser o mesmo ID do registro.")
			return errors
		}

		if !exists {
			errors["parent_id"] = append(errors["parent_id"], "A categoria pai informada não existe.")
			return errors
		}
	}

	logger.Info("---- Terminou de validar o payload da categoria ----")
	return errors
}
