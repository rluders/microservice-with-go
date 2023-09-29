package rest

import (
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"github.com/rluders/tutorial-microservices/menu-service/internal/domain"
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
				t.Fatalf("Failed to get path template: %v", err)
			}

			methods, err := route.GetMethods()
			if err != nil {
				t.Fatalf("Failed to get methods: %v", err)
			}

			for _, method := range methods {
				registeredRoutes[method+" "+pathTemplate] = true
			}
		}

		return nil
	}

	err := router.Walk(callback)
	if err != nil {
		t.Fatalf("An unexpected error happened while walking through the router: %v", err)
	}

	for _, route := range expectedRoutes {
		t.Run(route.path, func(t *testing.T) {
			expectedRouteKey := route.method + " " + route.path

			if _, ok := registeredRoutes[expectedRouteKey]; !ok {
				t.Errorf("Expected route %v is not registered", expectedRouteKey)
			}
		})
	}
}
