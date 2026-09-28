package financialmapper

import (
	financialresponse "finance/internal/http/response/financial"
	financialtransfermodel "finance/models/financial_transfer"
)

func ToFinancialTransferResponse(f financialtransfermodel.FinancialTransferModel) financialresponse.FinancialTransferResponse {
	return financialresponse.FinancialTransferResponse{
		ID:                   f.ID,
		TransferDate:         f.TransferDate,
		SourceAccountID:      f.SourceAccountID,
		DestinationAccountID: f.DestinationAccountID,
		IdempotencyKey:       f.IdempotencyKey,
		Amount:               f.Amount,
		CreatedAt:            f.CreatedAt,
		CanceledAt:           f.CanceledAt,
	}
}

func ToFinancialTransferResponseList(financialTransfers []financialtransfermodel.FinancialTransferModel) []financialresponse.FinancialTransferResponse {
	financialTransferResponse := make([]financialresponse.FinancialTransferResponse, 0, len(financialTransfers))

	for _, f := range financialTransfers {
		financialTransferResponse = append(financialTransferResponse, ToFinancialTransferResponse(f))
	}

	return financialTransferResponse
}
