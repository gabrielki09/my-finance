package routes

import (
	categoryesroutes "finance/internal/modules/categories/routes"
	"finance/internal/routes/cors"
	"log"
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

	categoryesroutes.RegisterCategoriesRoutes(publicMux, db)

	rootMux.Handle("/api/auth/", http.StripPrefix("/api/auth", publicMux))

	handlerWithCORS := cors.WithCORS(rootMux)

	log.Printf("Servidor rodando em http://localhost:%s/api", port)
	if err := http.ListenAndServe(":"+port, handlerWithCORS); err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}
