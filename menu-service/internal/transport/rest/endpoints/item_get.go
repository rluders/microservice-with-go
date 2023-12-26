package endpoints

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rluders/tutorial-microservices/menu-service/internal/domain"
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
			sendErrorResponse(w, "ID not found in request", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			sendErrorResponse(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		request := &GetItemRequest{ID: id}

		if err := isRequestValid(request); err != nil {
			sendValidationError(w, err)
			return
		}

		item, err := itemService.Get(request.ID)
		if err != nil {
			sendErrorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		payload := &GetItemResponse{
			Item: item,
		}
		sendDataResponse[GetItemResponse](w, "Item found", http.StatusOK, payload)
	}
}
