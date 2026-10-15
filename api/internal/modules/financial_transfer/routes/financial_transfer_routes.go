package financialtransferroutes

import (
	"net/http"

	financialtransfercontroller "finance/internal/modules/financial_transfer/controller"
	financialtransferrepository "finance/internal/modules/financial_transfer/repository"
	financialtransferservice "finance/internal/modules/financial_transfer/service"

	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterFinancialTransferRoutes(r *http.ServeMux, db *pgxpool.Pool) {
	repo := financialtransferrepository.NewFinancialTransferRepository(db)
	service := financialtransferservice.NewFinancialTransferService(repo)
	controller := financialtransfercontroller.NewFinancialTransferController(service)

	r.HandleFunc("GET /financial-transfer", controller.GetAll)
	r.HandleFunc("GET /financial-transfer/{id}", controller.FindByID)
	r.HandleFunc("POST /financial-transfer", controller.Create)
	r.HandleFunc("DELETE /financial-transfer/cancel/{id}", controller.Delete)
}
