package domain

type ItemRepository interface {
	CreateItem(item *Item) error
	UpdateItem(item *Item) error
	DeleteItem(itemID int) error
	FindItemByID(itemID int) (*Item, error)
	ListItems() ([]*Item, error)
}
