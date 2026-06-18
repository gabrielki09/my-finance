package financialaccountroutes

import (
	financialaccountcontroller "finance/internal/modules/financial_accounts/controller"
	financialaccountrepository "finance/internal/modules/financial_accounts/repository"
	financialaccountservice "finance/internal/modules/financial_accounts/services"
	financialaccountvalidator "finance/internal/modules/financial_accounts/validator"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterFinancialAccountRoutes(r *http.ServeMux, db *pgxpool.Pool) {
	repo := financialaccountrepository.NewFinancialAccountRepository(db)
	validator := financialaccountvalidator.NewFinancialAccountValidatorValidator(repo)

	service := financialaccountservice.NewFinancialAccountService(repo, validator)
	controller := financialaccountcontroller.NewFinancialAccountController(service)

	r.HandleFunc("GET /financial-account", controller.GetAll)
	r.HandleFunc("GET /financial-account/{id}", controller.FindById)
	r.HandleFunc("POST /financial-account", controller.Create)
	r.HandleFunc("PUT /financial-account/{id}", controller.Update)
	r.HandleFunc("DELETE /financial-account/delete/{id}", controller.Delete)
	r.HandleFunc("PATCH /financial-account/active/{id}", controller.Active)
	r.HandleFunc("GET /financial-account/current-balance/{id}", controller.GetCurrentBalance)
}
