package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	"menu-service/internal/domain"
	"menu-service/internal/transport/rest/endpoints"
)

type ItemHandler struct {
	itemService *domain.ItemService
}

func NewItemHandler(itemService *domain.ItemService) *ItemHandler {
	return &ItemHandler{
		itemService: itemService,
	}
}

func (h *ItemHandler) Register(router *mux.Router) {
	listItemEndpoint := endpoints.MakeListItemEndpoint(h.itemService)
	getItemEndpoint := endpoints.MakeGetItemEndpoint(h.itemService)

	router.HandleFunc("/items", listItemEndpoint).Methods(http.MethodGet)
	router.HandleFunc("/items/{id}", getItemEndpoint).Methods(http.MethodGet)

	protected := router.PathPrefix("/").Subrouter()
	//protected.Use(AuthMiddleware)

	createItemEndpoint := endpoints.MakeCreateItemEndpoint(h.itemService)
	updateItemEndpoint := endpoints.MakeUpdateItemEndpoint(h.itemService)
	deleteItemEndpoint := endpoints.MakeDeleteItemEndpoint(h.itemService)

	protected.HandleFunc("/items", createItemEndpoint).Methods(http.MethodPost)
	protected.HandleFunc("/items/{id}", updateItemEndpoint).Methods(http.MethodPut)
	protected.HandleFunc("/items/{id}", deleteItemEndpoint).Methods(http.MethodDelete)
}
