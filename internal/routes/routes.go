package routes

import (
	"finance/internal/helpers/response"
	"finance/internal/logger"
	categoryesroutes "finance/internal/modules/categories/routes"
	financialtransactionroutes "finance/internal/modules/financial/routes"
	financialaccountroutes "finance/internal/modules/financial_accounts/routes"
	financialobligationsroutes "finance/internal/modules/financial_obligations/routes"
	"finance/internal/routes/cors"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func StartServer(db *pgxpool.Pool) {
	rootMux := http.NewServeMux()
	publicMux := http.NewServeMux()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

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

	logger.General.Info.Printf("Servidor rodando em http://localhost:%s/api", port)

	if err := http.ListenAndServe(":"+port, handlerWithCORS); err != nil {
		logger.General.Info.Fatal("Erro ao iniciar o servidor:", err)
	}
}
