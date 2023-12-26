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
			sendErrorResponse(w, "ID not found in request", http.StatusBadRequest)
			return
		}

		categoryID, err := strconv.Atoi(idStr)
		if err != nil {
			sendErrorResponse(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		request := &UpdateCategoryRequest{
			ID: categoryID,
		}

		if err := parseRequest(request, r.Body); err != nil {
			sendErrorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := isRequestValid(request); err != nil {
			sendValidationError(w, err)
			return
		}

		category := &domain.Category{ID: request.ID, Name: request.Name}
		if err := categoryService.Update(category); err != nil {
			sendErrorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		payload := &UpdateCategoryResponse{
			Category: category,
		}
		sendDataResponse[UpdateCategoryResponse](w, "Category updated", http.StatusOK, payload)
	}
}
