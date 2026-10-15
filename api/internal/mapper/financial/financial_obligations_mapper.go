package financialmapper

import (
	financialresponse "finance/internal/http/response/financial"
	financialmodel "finance/models/financial"
)

func ToFinancialObligationResponse(f financialmodel.FinancialObligationModel) financialresponse.FinancialObligationResponse {
	return financialresponse.FinancialObligationResponse{
		Id:             f.Id,
		CategoryId:     f.CategoryId,
		Description:    f.Description,
		Type:           f.Type,
		Status:         f.Status,
		OriginalAmount: f.OriginalAmount,
		DueDate:        f.DueDate,
		CompetenceDate: f.CompetenceDate,
		Notes:          f.Notes,
		CreatedAt:      f.CreatedAt,
		UpdatedAt:      f.UpdatedAt,
	}
}

func ToFinancialObligationResponseList(financialObligations []financialmodel.FinancialObligationModel) []financialresponse.FinancialObligationResponse {
	financialObligationResponse := make([]financialresponse.FinancialObligationResponse, 0, len(financialObligations))

	for _, f := range financialObligations {
		financialObligationResponse = append(financialObligationResponse, ToFinancialObligationResponse(f))
	}

	return financialObligationResponse
}
