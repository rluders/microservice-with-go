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
			sendResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		if len(categories) == 0 {
			sendResponse(w, "Categories not found", http.StatusNotFound)
			return
		}

		payload := &ListCategoryResponse{
			Categories: categories,
		}
		sendDataResponse[ListCategoryResponse](w, "Categories found", http.StatusOK, payload)
	}
}
