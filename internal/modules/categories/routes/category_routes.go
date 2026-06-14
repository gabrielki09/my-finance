package categoryesroutes

import (
	categoriescontroller "finance/internal/modules/categories/controller"
	categoriesrepository "finance/internal/modules/categories/repository"
	categoriesservice "finance/internal/modules/categories/services"
	categoryvalidator "finance/internal/modules/categories/validator"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterCategoriesRoutes(r *http.ServeMux, db *pgxpool.Pool) {
	repo := categoriesrepository.NewCategoryRepository(db)
	categoryValidator := categoryvalidator.NewCategoryValidator(repo)

	service := categoriesservice.NewCategoryService(repo, categoryValidator)
	controller := categoriescontroller.NewCategoryController(service)

	r.HandleFunc("GET /categories", controller.GetAll)
	r.HandleFunc("GET /categories/{id}", controller.FindById)
	r.HandleFunc("POST /categories", controller.Create)
	r.HandleFunc("PUT /categories/{id}", controller.Update)
	r.HandleFunc("DELETE /categories/delete/{id}", controller.Delete)
	r.HandleFunc("PATCH /categories/active/{id}", controller.Active)
}
