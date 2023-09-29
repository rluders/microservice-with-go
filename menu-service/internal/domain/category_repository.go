package domain

type CategoryRepository interface {
	CreateCategory(category *Category) error
	UpdateCategory(category *Category) error
	DeleteCategory(categoryID int) error
	FindCategoryByID(categoryID int) (*Category, error)
	ListCategories() ([]*Category, error)
	// handle items
	AddItemToCategory(itemID, categoryID int) error
	RemoveItemFromCategory(itemID, categoryID int) error
	ItemCategories(itemID int) ([]*Category, error)
}
