package domain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItemJSONSerialization(t *testing.T) {
	// Create an example item
	item := &Item{
		ID:          1,
		Name:        "Test Item",
		Description: "Description of test item",
		Price:       9.99,
	}

	// Serialize the item to JSON
	data, err := json.Marshal(item)
	assert.NoError(t, err, "Error serializing item to JSON")

	// Deserialize JSON back to an item
	var newItem Item
	err = json.Unmarshal(data, &newItem)
	assert.NoError(t, err, "Error deserializing JSON to item")

	// Use assert to compare the original item with the deserialized item
	assert.Equal(t, item, &newItem, "Original item and deserialized item do not match")

	// Use assert to compare the categories of the original item with the deserialized item
	assert.Len(t, newItem.Categories, len(item.Categories), "Number of categories does not match")
	for i := range item.Categories {
		assert.Equal(t, item.Categories[i].Name, newItem.Categories[i].Name, "Category name does not match")
	}
}
