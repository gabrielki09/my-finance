package financialobligationsroutes

import (
	financialobligationrequest "finance/internal/http/request/financial/financial_obligation"
	financialobligationcontroller "finance/internal/modules/financial_obligations/controller"
	financialobligationrepository "finance/internal/modules/financial_obligations/repository"
	financialobligationservice "finance/internal/modules/financial_obligations/service"

	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterFinancialObligationRoutes(r *http.ServeMux, db *pgxpool.Pool) {
	repo := financialobligationrepository.NewFinancialObligationRepository(db)
	requet := financialobligationrequest.NewFinancialObligationRequest(repo)
	service := financialobligationservice.NewFinancialObligationService(repo, requet)
	controller := financialobligationcontroller.NewFinancialObligationController(service)

	r.HandleFunc("GET /financial-obligation", controller.GetAll)
	r.HandleFunc("POST /financial-obligation", controller.Create)
	r.HandleFunc("PUT /financial-obligation/{id}", controller.Update)
	r.HandleFunc("DELETE /financial-obligation/{id}", controller.Cancel)
}
