package endpoints

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"menu-service/internal/domain"
)

func TestMakeUpdateItemEndpoint(t *testing.T) {
	// Create a mock instance of ItemService to pass to the function.
	itemService := &domain.ItemService{}

	// Create a fake HTTP server to handle the request.
	server := httptest.NewServer(MakeUpdateItemEndpoint(itemService))
	defer server.Close()

	// Create a fake PUT request to the fake server.
	req, err := http.NewRequest("PUT", server.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Perform the request and get the response.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(resp.Body)

	// Assert that the response status code is http.StatusInternalServerError (500).
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}
