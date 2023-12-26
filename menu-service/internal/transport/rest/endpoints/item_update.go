package endpoints

import (
	"net/http"

	"menu-service/internal/domain"
)

func MakeUpdateItemEndpoint(itemService *domain.ItemService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "endpoint not implemented yet", http.StatusInternalServerError)
	}
}
