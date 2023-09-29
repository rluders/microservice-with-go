package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rluders/tutorial-microservices/menu-service/internal/domain"
	"github.com/rluders/tutorial-microservices/menu-service/internal/transport/rest/endpoints"
)

type CategoryHandler struct {
	categoryService *domain.CategoryService
}

func NewCategoryHandler(categoryService *domain.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}

func (h *CategoryHandler) Register(router *mux.Router) {
	listCategoryEndpoint := endpoints.MakeListCategoryEndpoint(h.categoryService)
	findCategoryEndpoint := endpoints.MakeFindCategoryEndpoint(h.categoryService)

	router.HandleFunc("/categories", listCategoryEndpoint).Methods(http.MethodGet)
	router.HandleFunc("/categories/{id}", findCategoryEndpoint).Methods(http.MethodGet)

	protected := router.PathPrefix("/").Subrouter()
	// protected.Use(AuthMiddleware)

	createCategoryEndpoint := endpoints.MakeCreateCategoryEndpoint(h.categoryService)
	updateCategoryEndpoint := endpoints.MakeUpdateCategoryEndpoint(h.categoryService)
	deleteCategoryEndpoint := endpoints.MakeDeleteCategoryEndpoint(h.categoryService)

	protected.HandleFunc("/categories", createCategoryEndpoint).Methods(http.MethodPost)
	protected.HandleFunc("/categories/{id}", updateCategoryEndpoint).Methods(http.MethodPut)
	protected.HandleFunc("/categories/{id}", deleteCategoryEndpoint).Methods(http.MethodDelete)
}
