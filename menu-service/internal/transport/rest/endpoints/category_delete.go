package endpoints

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rluders/tutorial-microservices/menu-service/internal/domain"
)

type DeleteCategoryRequest struct {
	ID int `json:"id" validate:"required"`
}

func MakeDeleteCategoryEndpoint(categoryService *domain.CategoryService) http.HandlerFunc {
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

		request := &DeleteCategoryRequest{ID: id}

		if err := isRequestValid(request); err != nil {
			sendValidationError(w, err)
			return
		}

		if err := categoryService.DeleteCategory(request.ID); err != nil {
			sendErrorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		sendResponse(w, "Category deleted", http.StatusOK)
	}
}