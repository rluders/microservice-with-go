package endpoints

import (
	"net/http"

	"github.com/rluders/tutorial-microservices/menu-service/internal/domain"
)

type ListItemResponse struct {
	Items []*domain.Item `json:"items"`
}

func MakeListItemEndpoint(itemService *domain.ItemService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items, err := itemService.ListItems()
		if err != nil {
			sendErrorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		payload := &ListItemResponse{
			Items: items,
		}
		sendDataResponse[ListItemResponse](w, "Items found", http.StatusOK, payload)
	}
}
