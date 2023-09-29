package endpoints

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rluders/tutorial-microservices/menu-service/internal/domain"
)

type FindCategoryByID struct {
	ID int `json:"id" validate:"required"`
}

type FindCategoryResponse struct {
	Category *domain.Category `json:"category"`
}

func MakeFindCategoryEndpoint(categoryService *domain.CategoryService) http.HandlerFunc {
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

		request := &FindCategoryByID{ID: id}

		if err := isRequestValid(request); err != nil {
			sendValidationError(w, err)
			return
		}

		category, err := categoryService.FindCategoryByID(request.ID)
		if err != nil {
			sendErrorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		payload := &FindCategoryResponse{
			Category: category,
		}
		sendDataResponse[FindCategoryResponse](w, "Category found", http.StatusOK, payload)
	}
}
