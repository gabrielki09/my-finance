package categoryesservice

import (
	"context"
	"finance/internal/apperrors"
	categoryresponse "finance/internal/http/response/category"
	"finance/internal/logger"
	categorymapper "finance/internal/mapper/category"
	categorymodel "finance/models/category"

	"github.com/google/uuid"
)

type CategoryRepository interface {
	GetAll(r context.Context) ([]categorymodel.CategoryModel, error)
	Create(r context.Context, category categorymodel.CategoryModel) (categorymodel.CategoryModel, error)
	Update(r context.Context, category categorymodel.CategoryModel) (categorymodel.CategoryModel, error)
	FindById(r context.Context, categoryId string) (categorymodel.CategoryModel, error)
	Delete(r context.Context, categoryId string) error
	Active(r context.Context, categoryId string) error
}

type CategoryService struct {
	repository CategoryRepository
}

func NewCategoryService(repository CategoryRepository) *CategoryService {
	return &CategoryService{
		repository: repository,
	}
}

func (s *CategoryService) GetAll(ctx context.Context) ([]categoryresponse.CategoryResponse, error) {
	categories, err := s.repository.GetAll(ctx)

	if err != nil {
		logger.General.Error.Println("Erro ao consultar todas as categorias:", err)
		return nil, err
	}

	return categorymapper.ToCategoryResponseList(categories), nil
}

func (s *CategoryService) Create(ctx context.Context, category categorymodel.CategoryModel) (categoryresponse.CategoryResponse, error) {
	category, err := s.repository.Create(ctx, category)

	if err != nil {
		logger.General.Error.Println("Erro ao criar a categoria:", err)
		return categoryresponse.CategoryResponse{}, err
	}

	return categorymapper.ToCategoryResponse(category), nil
}

func (s *CategoryService) Update(ctx context.Context, category categorymodel.CategoryModel) (categoryresponse.CategoryResponse, error) {
	category, err := s.repository.Update(ctx, category)

	if err != nil {
		logger.General.Error.Println("Erro ao alterar a categoria:", err)
		return categoryresponse.CategoryResponse{}, err
	}

	return categorymapper.ToCategoryResponse(category), nil
}

func (s *CategoryService) FindById(ctx context.Context, categoryId string) (categoryresponse.CategoryResponse, error) {

	if err := uuid.Validate(categoryId); err != nil {
		logger.General.Error.Println("UUID informado inválido.")
		return categoryresponse.CategoryResponse{}, apperrors.ErrInvalidUUID
	}

	category, err := s.repository.FindById(ctx, categoryId)

	if err != nil {
		logger.General.Error.Println("Erro ao localizar a categoria:", err)
		return categoryresponse.CategoryResponse{}, err
	}

	return categorymapper.ToCategoryResponse(category), nil
}

func (s *CategoryService) Delete(ctx context.Context, categoryId string) error {
	if err := uuid.Validate(categoryId); err != nil {
		logger.General.Error.Println("UUID informado inválido.")
		return apperrors.ErrInvalidUUID
	}

	if err := s.repository.Delete(ctx, categoryId); err != nil {
		logger.General.Error.Println("Erro ao deletar a categoria:", err)
		return err
	}

	return nil
}

func (s *CategoryService) Active(ctx context.Context, categoryId string) error {
	if err := uuid.Validate(categoryId); err != nil {
		logger.General.Error.Println("UUID informado inválido.")
		return apperrors.ErrInvalidUUID
	}

	if err := s.repository.Active(ctx, categoryId); err != nil {
		logger.General.Error.Println("Erro ao ativar a categoria:", err)
		return err
	}

	return nil
}
