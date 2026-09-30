package categoryesservice

import (
	"context"
	"finance/internal/apperrors"
	categoryrequest "finance/internal/http/request_v2/category"
	categoryresponse "finance/internal/http/response/category"
	"finance/internal/logger"
	categorymapper "finance/internal/mapper/category"
	categorymodel "finance/models/category"
)

type CategoryRepository interface {
	GetAll(r context.Context) ([]categorymodel.CategoryModel, error)
	Create(r context.Context, category categoryrequest.CategoryRequest) (categorymodel.CategoryModel, error)
	Update(r context.Context, category categoryrequest.CategoryRequest, categoryId int) (categorymodel.CategoryModel, error)
	FindById(r context.Context, categoryId int) (categorymodel.CategoryModel, error)
	Delete(r context.Context, categoryId int) error
	Active(r context.Context, categoryId int) error
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
		logger.Error("Erro ao consultar todas as categorias:", err)
		return nil, err
	}

	return categorymapper.ToCategoryResponseList(categories), nil
}

func (s *CategoryService) Create(ctx context.Context, payload categoryrequest.CategoryRequest) (categoryresponse.CategoryResponse, error) {
	validation := payload.ValidatePayload(ctx)

	if len(validation) > 0 {
		return categoryresponse.CategoryResponse{}, apperrors.NewValidationError(validation)
	}

	category, err := s.repository.Create(ctx, payload)

	if err != nil {
		logger.Error("Erro ao criar a categoria:", err)
		return categoryresponse.CategoryResponse{}, err
	}

	return categorymapper.ToCategoryResponse(category), nil
}

func (s *CategoryService) Update(ctx context.Context, payload categoryrequest.CategoryRequest, categoryId int) (categoryresponse.CategoryResponse, error) {

	payload.Id = categoryId
	validation := payload.ValidatePayload(ctx)

	if len(validation) > 0 {
		return categoryresponse.CategoryResponse{}, apperrors.NewValidationError(validation)
	}

	category, err := s.repository.Update(ctx, payload, categoryId)

	if err != nil {
		logger.Error("Erro ao alterar a categoria:", err)
		return categoryresponse.CategoryResponse{}, err
	}

	return categorymapper.ToCategoryResponse(category), nil
}

func (s *CategoryService) FindById(ctx context.Context, categoryId int) (categoryresponse.CategoryResponse, error) {
	category, err := s.repository.FindById(ctx, categoryId)

	if err != nil {
		logger.Error("Erro ao localizar a categoria:", err)
		return categoryresponse.CategoryResponse{}, err
	}

	return categorymapper.ToCategoryResponse(category), nil
}

func (s *CategoryService) Delete(ctx context.Context, categoryId int) error {
	if err := s.repository.Delete(ctx, categoryId); err != nil {
		logger.Error("Erro ao deletar a categoria:", err)
		return err
	}

	return nil
}

func (s *CategoryService) Active(ctx context.Context, categoryId int) error {
	if err := s.repository.Active(ctx, categoryId); err != nil {
		logger.Error("Erro ao ativar a categoria:", err)
		return err
	}

	return nil
}
