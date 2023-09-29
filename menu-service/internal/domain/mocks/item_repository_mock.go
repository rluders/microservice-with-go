package mocks

import "github.com/rluders/tutorial-microservices/menu-service/internal/domain"

type ItemRepositoryMock struct {
	data map[int]*domain.Item
}

func NewItemRepositoryMock(data map[int]*domain.Item) *ItemRepositoryMock {
	return &ItemRepositoryMock{
		data: make(map[int]*domain.Item),
	}
}

func (r *ItemRepositoryMock) CreateItem(item *domain.Item) error {
	if item.ID == 0 {
		item.ID = len(r.data) + 1
	}
	r.data[item.ID] = item
	return nil
}

func (r *ItemRepositoryMock) UpdateItem(item *domain.Item) error {
	if _, ok := r.data[item.ID]; !ok {
		return domain.ErrItemNotFound
	}
	r.data[item.ID] = item
	return nil
}

func (r *ItemRepositoryMock) DeleteItem(itemID int) error {
	if _, ok := r.data[itemID]; !ok {
		return domain.ErrItemNotFound
	}
	delete(r.data, itemID)
	return nil
}

func (r *ItemRepositoryMock) FindItemByID(itemID int) (*domain.Item, error) {
	item, ok := r.data[itemID]
	if !ok {
		return nil, domain.ErrItemNotFound
	}
	return item, nil
}

func (r *ItemRepositoryMock) ListItems() ([]*domain.Item, error) {
	var items []*domain.Item
	for _, item := range r.data {
		items = append(items, item)
	}
	return items, nil
}
