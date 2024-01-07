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
			sendResponse(w, "ID not found in request", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			sendResponse(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		request := &GetItemRequest{ID: id}

		if err := isRequestValid(request); err != nil {
			sendValidationError(w, err)
			return
		}

		item, err := itemService.Get(request.ID)
		if err != nil {
			sendResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		payload := &GetItemResponse{
			Item: item,
		}
		sendDataResponse[GetItemResponse](w, "Item found", http.StatusOK, payload)
	}
}
