package financialobligationsroutes

import (
	financialobligationcontroller "finance/internal/modules/financial_obligations/controller"
	financialobligationrepository "finance/internal/modules/financial_obligations/repository"
	financialobligationservice "finance/internal/modules/financial_obligations/service"
	financialobligationvalidator "finance/internal/modules/financial_obligations/validator"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterFinancialObligationRoutes(r *http.ServeMux, db *pgxpool.Pool) {
	repo := financialobligationrepository.NewFinancialObligationRepository(db)
	validator := financialobligationvalidator.NewFinancialObligationValidatorValidator(repo)
	service := financialobligationservice.NewFinancialObligationService(repo, validator)
	controller := financialobligationcontroller.NewFinancialObligationController(service)

	r.HandleFunc("GET /financial-obligation", controller.GetAll)
	r.HandleFunc("POST /financial-obligation", controller.Create)
	r.HandleFunc("PUT /financial-obligation/{id}", controller.Update)
	r.HandleFunc("DELETE /financial-obligation/cancel/{id}", controller.Cancel)
	r.HandleFunc("POST /financial-obligation/pay", controller.Pay)

}
