package endpoints

import (
	"net/http"

	"menu-service/internal/domain"
)

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

type CreateCategoryResponse struct {
	Category *domain.Category `json:"category"`
}

func MakeCreateCategoryEndpoint(categoryService *domain.CategoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := &CreateCategoryRequest{}

		if err := parseRequest(request, r.Body); err != nil {
			sendErrorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := isRequestValid(request); err != nil {
			sendValidationError(w, err)
			return
		}

		category := &domain.Category{Name: request.Name}
		if err := categoryService.Create(category); err != nil {
			sendErrorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		payload := &CreateCategoryResponse{
			Category: category,
		}
		sendDataResponse[CreateCategoryResponse](w, "Category created", http.StatusCreated, payload)
	}
}
