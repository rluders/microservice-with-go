package domain

// ItemRepository represents the repository interface for managing items in a data store.
type ItemRepository interface {
	// Create adds a new item to the data store.
	Create(item *Item) error
	// Update modifies an existing item in the data store.
	Update(item *Item) error
	// Delete removes an item from the data store based on its ID.
	Delete(itemID int) error
	// Get retrieves an item from the data store based on its ID.
	Get(itemID int) (*Item, error)
	// List retrieves a list of all items from the data store.
	List() ([]*Item, error)
}
