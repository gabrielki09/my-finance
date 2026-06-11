package categorymapper

import (
	categoryresponse "finance/internal/http/response/category"
	categorymodel "finance/models/category"
)

func ToCategoryResponse(c categorymodel.CategoryModel) categoryresponse.CategoryResponse {
	return categoryresponse.CategoryResponse{
		Id:   c.Id,
		Name: c.Name,
	}
}

func ToCategoryResponseList(categoriesList []categorymodel.CategoryModel) []categoryresponse.CategoryResponse {
	categoriesResponse := make([]categoryresponse.CategoryResponse, 0, len(categoriesList))

	for _, c := range categoriesList {
		categoriesResponse = append(categoriesResponse, ToCategoryResponse(c))
	}

	return categoriesResponse
}
