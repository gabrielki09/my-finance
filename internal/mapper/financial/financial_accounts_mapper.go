package financialmapper

import (
	financialresponse "finance/internal/http/response/financial"
	financialmodel "finance/models/financial"
)

func ToFinancialAccountResponse(f financialmodel.FinancialAccountModel) financialresponse.FinancialAccountResponse {
	return financialresponse.FinancialAccountResponse{
		Id:             f.Id,
		Name:           f.Name,
		Type:           f.Type,
		InitialBalance: f.InitialBalance,
		OpenedAt:       f.OpenedAt,
		CreatedAt:      f.CreatedAt,
		UpdatedAt:      f.UpdatedAt,
	}
}

func ToFinancialAccountResponseList(financialAccount []financialmodel.FinancialAccountModel) []financialresponse.FinancialAccountResponse {
	financialAccountResponse := make([]financialresponse.FinancialAccountResponse, 0, len(financialAccount))

	for _, f := range financialAccount {
		financialAccountResponse = append(financialAccountResponse, ToFinancialAccountResponse(f))
	}

	return financialAccountResponse
}
