package endpoints

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"menu-service/internal/domain"
)

type DeleteCategoryRequest struct {
	ID int `json:"id" validate:"required"`
}

func MakeDeleteCategoryEndpoint(categoryService *domain.CategoryService) http.HandlerFunc {
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

		request := &DeleteCategoryRequest{ID: id}
		if err := isRequestValid(request); err != nil {
			sendResponse[ValidationErrors](w, "Validation error", http.StatusBadRequest, nil)
			return
		}

		if err := categoryService.Delete(request.ID); err != nil {
			sendResponse[any](w, err.Error(), http.StatusBadRequest, nil)
			return
		}

		sendResponse[any](w, "Category deleted", http.StatusOK, nil)
	}
}
