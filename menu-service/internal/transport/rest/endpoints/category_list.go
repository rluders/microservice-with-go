package endpoints

import (
	"net/http"

	"github.com/rluders/tutorial-microservices/menu-service/internal/domain"
)

type ListCategoryResponse struct {
	Categories []*domain.Category `json:"categories"`
}

func MakeListCategoryEndpoint(categoryService *domain.CategoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		categories, err := categoryService.ListCategories()
		if err != nil {
			sendErrorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		payload := &ListCategoryResponse{
			Categories: categories,
		}
		sendDataResponse[ListCategoryResponse](w, "Categories found", http.StatusOK, payload)
	}
}
