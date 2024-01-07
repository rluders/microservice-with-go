package endpoints

import (
	"net/http"

	"menu-service/internal/domain"
)

type CreateItemRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Categories  []int   `json:"categories"`
}

type CreateItemResponse struct {
	Item *domain.Item `json:"item"`
}

func MakeCreateItemEndpoint(itemService *domain.ItemService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := &CreateItemRequest{}

		if err := parseRequest(request, r.Body); err != nil {
			sendResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := isRequestValid(request); err != nil {
			sendValidationError(w, err)
			return
		}

		item := &domain.Item{
			Name:        request.Name,
			Description: request.Description,
			Price:       request.Price,
		}
		//if len(request.Categories) > 0 {
		//	item.Categories = []*domain.Category{}
		//	for _, c := range request.Categories {
		//		item.Categories = append(item.Categories, &domain.Category{ID: c})
		//	}
		//}

		if err := itemService.Create(item); err != nil {
			sendResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		payload := &CreateItemResponse{
			Item: item,
		}
		sendDataResponse[CreateItemResponse](w, "Item created", http.StatusCreated, payload)
	}
}
