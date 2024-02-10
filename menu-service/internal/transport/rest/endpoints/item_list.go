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
			sendResponse[any](w, err.Error(), http.StatusBadRequest, nil)
			return
		}

		if len(items) == 0 {
			sendResponse[any](w, "Items not found", http.StatusNotFound, nil)
			return
		}

		body := &ListItemResponse{
			Items: items,
		}
		sendResponse[ListItemResponse](w, "Items found", http.StatusOK, body)
	}
}
