package endpoints

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"menu-service/internal/domain"
)

type GetItemRequest struct {
	ID int `json:"id" validate:"required"`
}

type GetItemResponse struct {
	Item *domain.Item `json:"item"`
}

func MakeGetItemEndpoint(itemService *domain.ItemService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr, ok := vars["id"]
		if !ok {
			sendResponse[any](w, "ID not found in request", http.StatusBadRequest, nil)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			sendResponse[any](w, "Invalid ID format", http.StatusBadRequest, nil)
			return
		}

		request := &GetItemRequest{ID: id}
		if err := isRequestValid(request); err != nil {
			sendResponse[ValidationErrors](w, "Validation error", http.StatusBadRequest, err)
			return
		}

		item, err := itemService.Get(request.ID)
		if err != nil {
			sendResponse[any](w, err.Error(), http.StatusBadRequest, nil)
			return
		}

		body := &GetItemResponse{
			Item: item,
		}
		sendResponse[GetItemResponse](w, "Item found", http.StatusOK, body)
	}
}
