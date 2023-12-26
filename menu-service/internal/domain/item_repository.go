package domain

type ItemRepository interface {
	Create(item *Item) error
	Update(item *Item) error
	Delete(itemID int) error
	Get(itemID int) (*Item, error)
	List() ([]*Item, error)
}
