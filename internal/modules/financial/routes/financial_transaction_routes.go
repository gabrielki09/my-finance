package financialtransactionroutes

import (
	financialtransactioncontroller "finance/internal/modules/financial/controller"
	financialtransactionrepository "finance/internal/modules/financial/repository"
	financialtransactionservice "finance/internal/modules/financial/services"
	financialtransactionvalidator "finance/internal/modules/financial/validator"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterFinancialTransactionRoutes(r *http.ServeMux, db *pgxpool.Pool) {
	repo := financialtransactionrepository.NewFinancialTransactionRepository(db)
	service := financialtransactionservice.NewFinancialTransactionService(repo, financialtransactionvalidator.NewFinancialTransactionValidator(repo))
	controller := financialtransactioncontroller.NewFinancialTransactionController(service)

	r.HandleFunc("GET /financial-transaction", controller.GetAll)
	r.HandleFunc("GET /financial-transaction/{id}", controller.FindById)
	r.HandleFunc("GET /financial-transaction/key/{key}", controller.FindByKey)
	r.HandleFunc("POST /financial-transaction", controller.CreateMovement)
	//r.HandleFunc("POST /financial-transaction", controller.CancelMovement)

}
