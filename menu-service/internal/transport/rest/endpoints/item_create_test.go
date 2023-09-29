package endpoints

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rluders/tutorial-microservices/menu-service/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestMakeCreateItemEndpoint(t *testing.T) {
	// Create a mock instance of ItemService to pass to the function.
	itemService := &domain.ItemService{}

	// Create a fake HTTP server to handle the request.
	server := httptest.NewServer(MakeCreateItemEndpoint(itemService))
	defer server.Close()

	// Create a fake POST request to the fake server.
	req, err := http.NewRequest("POST", server.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Perform the request and get the response.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Assert that the response status code is http.StatusInternalServerError (500).
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}
