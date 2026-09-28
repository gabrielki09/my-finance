package financialmapper

import (
	financialresponse "finance/internal/http/response/financial"
	financialmodel "finance/models/financial"
)

func ToFinancialTransactionResponse(f financialmodel.FinancialTransactionsModel) financialresponse.FinancialTransactionsResponse {
	return financialresponse.FinancialTransactionsResponse{
		Id:                    f.Id,
		FinancialAccountId:    f.FinancialAccountId,
		CategoryId:            f.CategoryId,
		ReversedTransactionId: f.ReversedTransactionId,
		Description:           f.Description,
		MovementType:          f.MovementType,
		OperationType:         f.OperationType,
		Amount:                f.Amount,
		MovementDate:          f.MovementDate,
		ReferenceDate:         f.ReferenceDate,
		OriginType:            f.OriginType,
		OriginId:              f.OriginId,
		IdempotencyKey:        f.IdempotencyKey,
		CreatedAt:             f.CreatedAt,
		CanceledAt:            f.CanceledAt,
	}
}

func ToFinancialTransactionResponseList(financialTransaction []financialmodel.FinancialTransactionsModel) []financialresponse.FinancialTransactionsResponse {
	financialTransactionResponse := make([]financialresponse.FinancialTransactionsResponse, 0, len(financialTransaction))

	for _, f := range financialTransaction {
		financialTransactionResponse = append(financialTransactionResponse, ToFinancialTransactionResponse(f))
	}

	return financialTransactionResponse
}
