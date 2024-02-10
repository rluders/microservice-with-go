package domain

// CategoryRepository represents the repository interface for managing categories in a data store.
type CategoryRepository interface {
	// Create adds a new category to the data store.
	Create(category *Category) error
	// Update modifies an existing category in the data store.
	Update(category *Category) error
	// Delete removes a category from the data store based on its ID.
	Delete(categoryID int) error
	// Get retrieves a category from the data store based on its ID.
	Get(categoryID int) (*Category, error)
	// List retrieves a list of all categories from the data store.
	List() ([]*Category, error)
	// AddItem associates an item with a category in the data store.
	AddItem(itemID, categoryID int) error
	// RemoveItem disassociates an item from a category in the data store.
	RemoveItem(itemID, categoryID int) error
}
