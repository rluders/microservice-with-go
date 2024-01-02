package endpoints

import (
	"net/http"

	"menu-service/internal/domain"
)

type ListItemResponse struct {
	Items []*domain.Item `json:"items"`
}

func MakeListItemEndpoint(itemService *domain.ItemService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items, err := itemService.List()
		if err != nil {
			sendErrorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		if len(items) == 0 {
			sendErrorResponse(w, "Items not found", http.StatusNotFound)
			return
		}

		payload := &ListItemResponse{
			Items: items,
		}
		sendDataResponse[ListItemResponse](w, "Items found", http.StatusOK, payload)
	}
}
