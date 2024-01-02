package rest

import (
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"menu-service/internal/domain"
)

func TestCategoryHandler_Routes(t *testing.T) {
	mockCategoryService := &domain.CategoryService{}

	router := mux.NewRouter()
	categoryHandler := NewCategoryHandler(mockCategoryService)
	categoryHandler.Register(router)

	expectedRoutes := []struct {
		method string
		path   string
	}{
		{http.MethodGet, "/categories"},
		{http.MethodGet, "/categories/{id}"},
		{http.MethodPost, "/categories"},
		{http.MethodPut, "/categories/{id}"},
		{http.MethodDelete, "/categories/{id}"},
	}

	registeredRoutes := make(map[string]bool)

	callback := func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		if methods, _ := route.GetMethods(); methods != nil {
			pathTemplate, err := route.GetPathTemplate()
			if err != nil {
				return err
			}

			methods, err := route.GetMethods()
			if err != nil {
				return err
			}

			for _, method := range methods {
				registeredRoutes[method+" "+pathTemplate] = true
			}
		}

		return nil
	}

	err := router.Walk(callback)
	assert.Nil(t, err, "Unexpected error while walking through the router")

	for _, r := range expectedRoutes {
		route := r
		t.Run(route.path, func(t *testing.T) {
			t.Parallel()
			expectedRouteKey := route.method + " " + route.path

			assert.True(t, registeredRoutes[expectedRouteKey], "Expected route %v is not registered", expectedRouteKey)
		})
	}
}
