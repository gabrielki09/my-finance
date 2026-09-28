package categoryvalidator

import (
	"context"
	"finance/internal/apperrors"
	categoryrequest "finance/internal/http/request/category"
	"finance/internal/logger"
	categorymodel "finance/models/category"
	"fmt"
)

type CategoryRepository interface {
	VerifyParentId(ctx context.Context, parentId int) (bool, error)
	VerifyExistsCategoryName(ctx context.Context, categoryName string) (*categorymodel.CategoryModel, error)
}
type CategoryValidator struct {
	repo CategoryRepository
}

func NewCategoryValidator(categoryRepository CategoryRepository) *CategoryValidator {
	return &CategoryValidator{
		repo: categoryRepository,
	}
}

func (v *CategoryValidator) ValidatePayload(ctx context.Context, payload categoryrequest.CategoryRequest) error {
	logger.Info("---- Vai validar o payload da categoria via db ----")

	errors := apperrors.ValidationErrors{}

	if payload.ParentId != nil {
		exists, err := v.repo.VerifyParentId(ctx, *payload.ParentId)

		if err != nil {
			logger.Error("Erro ao conferir se a categoria pai existe: ", err)
			return err
		}

		if !exists {
			errors["parent_id"] = append(errors["parent_id"], "A categoria pai informada não existe.")
		}
	}

	categoryByName, err := v.repo.VerifyExistsCategoryName(ctx, payload.Name)

	if err != nil {
		logger.Error("Erro ao conferir se a categoria já existe pelo nome: ", err)
		return err
	}

	if categoryByName != nil {
		errors["name"] = append(errors["name"], fmt.Sprintf("A categoria %s já existe, ID %d.", payload.Name, categoryByName.Id))
	}

	logger.Info("---- Terminou de validar o payload da categoria, total de erros: %d ----", len(errors))

	if len(errors) > 0 {
		return apperrors.NewValidationError(errors)
	}

	return nil
}
