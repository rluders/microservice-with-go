package endpoints

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"menu-service/internal/domain"
)

type UpdateCategoryRequest struct {
	ID   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type UpdateCategoryResponse struct {
	Category *domain.Category `json:"category"`
}

func MakeUpdateCategoryEndpoint(categoryService *domain.CategoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr, ok := vars["id"]
		if !ok {
			sendResponse[any](w, "ID not found in request", http.StatusBadRequest, nil)
			return
		}

		categoryID, err := strconv.Atoi(idStr)
		if err != nil {
			sendResponse[any](w, "Invalid ID format", http.StatusBadRequest, nil)
			return
		}

		request := &UpdateCategoryRequest{
			ID: categoryID,
		}
		if err := parseRequest(request, r.Body); err != nil {
			sendResponse[any](w, err.Error(), http.StatusBadRequest, nil)
			return
		}

		if err := isRequestValid(request); err != nil {
			sendResponse[ValidationErrors](w, "Validation error", http.StatusBadRequest, err)
			return
		}

		category := &domain.Category{ID: request.ID, Name: request.Name}
		if err := categoryService.Update(category); err != nil {
			sendResponse[any](w, err.Error(), http.StatusBadRequest, nil)
			return
		}

		body := &UpdateCategoryResponse{
			Category: category,
		}
		sendResponse[UpdateCategoryResponse](w, "Category updated", http.StatusOK, body)
	}
}
