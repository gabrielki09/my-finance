package categorycontroller

import (
	"context"
	"encoding/json"
	"finance/internal/helpers/response"
	categoryresponse "finance/internal/http/response/category"
	categorymodel "finance/models/category"
	"net/http"
)

type CategoryService interface {
	GetAll(context.Context) ([]categoryresponse.CategoryResponse, error)
	Create(context.Context, categorymodel.CategoryModel) (categoryresponse.CategoryResponse, error)
	Update(context.Context, categorymodel.CategoryModel) (categoryresponse.CategoryResponse, error)
	FindById(context.Context, string) (categoryresponse.CategoryResponse, error)
	Delete(context.Context, string) error
	Active(context.Context, string) error
}

type CategoryController struct {
	service CategoryService
}

func NewCategoryController(service CategoryService) *CategoryController {
	return &CategoryController{
		service: service,
	}
}

func (c *CategoryController) GetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := c.service.GetAll(r.Context())

	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao retornar todas as categorias",
			map[string]any{"error": err.Error()},
		))
		return
	}

	response.WriteJSON(w, http.StatusOK, response.SuccessResponse(
		"Todas as categorias",
		map[string]any{"categories": categories},
	))
}

func (c *CategoryController) Create(w http.ResponseWriter, r *http.Request) {
	var categoryPayLoad categorymodel.CategoryModel

	if err := json.NewDecoder(r.Body).Decode(&categoryPayLoad); err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao ler os dados",
			map[string]any{"error": err.Error()},
		))
		return
	}

	category, err := c.service.Create(r.Context(), categoryPayLoad)

	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao cadastrar a categoria",
			map[string]any{"error": err.Error()},
		))
		return
	}

	response.WriteJSON(w, http.StatusCreated, response.SuccessResponse(
		"Categoria cadastrada com sucesso!",
		map[string]any{"category": category},
	))
}

func (c *CategoryController) Update(w http.ResponseWriter, r *http.Request) {
	var categoryPayLoad categorymodel.CategoryModel

	if err := json.NewDecoder(r.Body).Decode(&categoryPayLoad); err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao ler os dados",
			map[string]any{"error": err.Error()},
		))
		return
	}

	categoryPayLoad.Id = r.PathValue("id")

	category, err := c.service.Update(r.Context(), categoryPayLoad)

	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao alterar a categoria",
			map[string]any{"error": err.Error()},
		))
		return
	}

	response.WriteJSON(w, http.StatusCreated, response.SuccessResponse(
		"Categoria alterada com sucesso!",
		map[string]any{"category": category},
	))
}

func (c *CategoryController) FindById(w http.ResponseWriter, r *http.Request) {
	category, err := c.service.FindById(r.Context(), r.PathValue("id"))

	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao localizar a categoria",
			map[string]any{"error": err.Error()},
		))
		return
	}

	response.WriteJSON(w, http.StatusOK, response.SuccessResponse(
		"Categoria localizada com sucesso!",
		map[string]any{"category": category},
	))
}

func (c *CategoryController) Delete(w http.ResponseWriter, r *http.Request) {
	if err := c.service.Delete(r.Context(), r.PathValue("id")); err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao desativar a categoria",
			map[string]any{"error": err.Error()},
		))
		return
	}

	response.WriteJSON(w, http.StatusOK, response.SuccessResponse(
		"Categoria desativada com sucesso!",
		map[string]any{},
	))
}

func (c *CategoryController) Active(w http.ResponseWriter, r *http.Request) {
	if err := c.service.Active(r.Context(), r.PathValue("id")); err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao ativar a categoria",
			map[string]any{"error": err.Error()},
		))
		return
	}

	response.WriteJSON(w, http.StatusOK, response.SuccessResponse(
		"Categoria ativada com sucesso!",
		map[string]any{},
	))
}
