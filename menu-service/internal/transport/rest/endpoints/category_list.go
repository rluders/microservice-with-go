package endpoints

import (
	"net/http"

	"menu-service/internal/domain"
)

type ListCategoryResponse struct {
	Categories []*domain.Category `json:"categories"`
}

func MakeListCategoryEndpoint(categoryService *domain.CategoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		categories, err := categoryService.List()
		if err != nil {
			sendResponse[any](w, err.Error(), http.StatusBadRequest, nil)
			return
		}

		if len(categories) == 0 {
			sendResponse[any](w, "Categories not found", http.StatusNotFound, nil)
			return
		}

		body := &ListCategoryResponse{
			Categories: categories,
		}
		sendResponse[ListCategoryResponse](w, "Categories found", http.StatusOK, body)
	}
}
