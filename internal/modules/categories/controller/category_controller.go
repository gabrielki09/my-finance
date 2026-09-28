package categorycontroller

import (
	"context"
	"encoding/json"
	"errors"
	"finance/internal/apperrors"
	"finance/internal/helpers/getidpath"
	"finance/internal/helpers/response"
	"finance/internal/http/httpx"
	categoryrequest "finance/internal/http/request/category"
	categoryresponse "finance/internal/http/response/category"
	"finance/internal/logger"
	"net/http"
)

type CategoryService interface {
	GetAll(context.Context) ([]categoryresponse.CategoryResponse, error)
	Create(context.Context, categoryrequest.CategoryRequest) (categoryresponse.CategoryResponse, error)
	Update(context.Context, categoryrequest.CategoryRequest, int) (categoryresponse.CategoryResponse, error)
	FindById(context.Context, int) (categoryresponse.CategoryResponse, error)
	Delete(context.Context, int) error
	Active(context.Context, int) error
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
	logger.Info("CategoryController - GetAll")

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
	var payload categoryrequest.CategoryRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&payload); err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao ler os dados",
			httpx.DecodeErrorMessage(err),
		))
		return
	}

	category, err := c.service.Create(r.Context(), payload)

	if err != nil {
		var validationErr *apperrors.ValidationError

		if errors.As(err, &validationErr) {
			response.WriteJSON(w, http.StatusUnprocessableEntity, response.ErrorResponse(
				"Erro de validação",
				validationErr.Errors,
			))
			return
		}

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
	var payLoad categoryrequest.CategoryRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&payLoad); err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao ler os dados",
			httpx.DecodeErrorMessage(err),
		))
		return
	}

	id, err := getidpath.GetIdPath(r)

	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao ler o identificador da categoria",
			map[string]any{"error": err.Error()},
		))
		return
	}

	category, err := c.service.Update(r.Context(), payLoad, id)

	if err != nil {
		var validationErr *apperrors.ValidationError

		if errors.As(err, &validationErr) {
			response.WriteJSON(w, http.StatusUnprocessableEntity, response.ErrorResponse(
				"Erro de validação",
				validationErr.Errors,
			))
			return
		}

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
	id, err := getidpath.GetIdPath(r)

	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao ler o identificador da categoria",
			map[string]any{"error": err.Error()},
		))
		return
	}

	category, err := c.service.FindById(r.Context(), id)

	if err != nil {
		if err == apperrors.ErrNotFound {
			response.WriteJSON(w, http.StatusNotFound, response.ErrorResponse(
				"Categoria não localizada",
				map[string]any{},
			))
			return
		}

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
	id, err := getidpath.GetIdPath(r)

	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao ler o identificador da conta financeira",
			map[string]any{"error": err.Error()},
		))
		return
	}

	if err := c.service.Delete(r.Context(), id); err != nil {
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
	id, err := getidpath.GetIdPath(r)

	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao ler o identificador da conta financeira",
			map[string]any{"error": err.Error()},
		))
		return
	}

	if err := c.service.Active(r.Context(), id); err != nil {
		response.WriteJSON(w, http.StatusBadRequest, response.ErrorResponse(
			"Erro ao ativar a categoria",
			map[string]any{"error": err.Error()},
		))
		return
	}

	response.WriteJSON(w, http.StatusOK, response.SuccessResponse(
		"Operação realizada com sucesso.",
		map[string]any{},
	))
}
