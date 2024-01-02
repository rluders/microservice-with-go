package domain

type CategoryRepository interface {
	Create(category *Category) error
	Update(category *Category) error
	Delete(categoryID int) error
	Get(categoryID int) (*Category, error)
	List() ([]*Category, error)
	AddItem(itemID, categoryID int) error
	RemoveItem(itemID, categoryID int) error
}
