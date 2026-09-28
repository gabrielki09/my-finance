package routes

import (
	"finance/internal/config"
	"finance/internal/helpers/response"
	"finance/internal/logger"
	categoryesroutes "finance/internal/modules/categories/routes"
	financialtransactionroutes "finance/internal/modules/financial/routes"
	financialaccountroutes "finance/internal/modules/financial_accounts/routes"
	financialobligationsroutes "finance/internal/modules/financial_obligations/routes"
	"finance/internal/routes/cors"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func StartServer(db *pgxpool.Pool, apiConfig *config.ApiConfig) {
	rootMux := http.NewServeMux()
	publicMux := http.NewServeMux()

	rootMux.HandleFunc("GET /api/health", func(w http.ResponseWriter, r *http.Request) {
		response.WriteJSON(w, http.StatusOK, response.SuccessResponse(
			"Api on",
			map[string]any{},
		))
	})

	categoryesroutes.RegisterCategoriesRoutes(publicMux, db)
	financialaccountroutes.RegisterFinancialAccountRoutes(publicMux, db)
	financialtransactionroutes.RegisterFinancialTransactionRoutes(publicMux, db)
	financialobligationsroutes.RegisterFinancialObligationRoutes(publicMux, db)

	rootMux.Handle("/api/", http.StripPrefix("/api", publicMux))

	handlerWithCORS := cors.WithCORS(rootMux)

	logger.Info(fmt.Sprintf("Servidor rodando em http://localhost:%s/api", apiConfig.Port))

	if err := http.ListenAndServe(":"+apiConfig.Port, handlerWithCORS); err != nil {
		logger.Error("Erro ao iniciar o servidor:", err)
	}
}
