package domain

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCategoryJSONSerialization(t *testing.T) {
	// Create an example category
	category := &Category{
		ID:        1,
		Name:      "Test Category",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Serialize the category to JSON
	data, err := json.Marshal(category)
	assert.NoError(t, err, "Error serializing category to JSON")

	// Deserialize JSON back to a category
	var newCategory Category
	err = json.Unmarshal(data, &newCategory)
	assert.NoError(t, err, "Error deserializing JSON to category")

	// Use assert to compare the original category with the deserialized category
	assert.Equal(t, category, &newCategory, "Original category and deserialized category do not match")
}
