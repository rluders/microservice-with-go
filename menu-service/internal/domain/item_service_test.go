package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockItemRepository é um mock para ItemRepository.
type MockItemRepository struct {
	mock.Mock
}

func (m *MockItemRepository) CreateItem(item *Item) error {
	args := m.Called(item)
	return args.Error(0)
}

func (m *MockItemRepository) UpdateItem(item *Item) error {
	args := m.Called(item)
	return args.Error(0)
}

func (m *MockItemRepository) DeleteItem(itemID int) error {
	args := m.Called(itemID)
	return args.Error(0)
}

func (m *MockItemRepository) FindItemByID(itemID int) (*Item, error) {
	args := m.Called(itemID)
	if item := args.Get(0); item != nil {
		return item.(*Item), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockItemRepository) ListItems() ([]*Item, error) {
	args := m.Called()
	return args.Get(0).([]*Item), args.Error(1)
}

func TestCreateItem(t *testing.T) {
	itemRepo := new(MockItemRepository)
	catetoryRepo := new(MockCategoryRepository)
	service := NewItemService(itemRepo, catetoryRepo)

	item := &Item{Name: "Item Test"}

	itemRepo.On("CreateItem", item).Return(nil)

	err := service.CreateItem(item)

	assert.NoError(t, err)
	itemRepo.AssertExpectations(t)
}

func TestUpdateItem(t *testing.T) {
	itemRepo := new(MockItemRepository)
	catetoryRepo := new(MockCategoryRepository)
	service := NewItemService(itemRepo, catetoryRepo)

	item := &Item{ID: 1, Name: "Item Test"}

	itemRepo.On("UpdateItem", item).Return(nil)

	err := service.UpdateItem(item)

	assert.NoError(t, err)
	itemRepo.AssertExpectations(t)
}

func TestDeleteItem(t *testing.T) {
	itemRepo := new(MockItemRepository)
	catetoryRepo := new(MockCategoryRepository)
	service := NewItemService(itemRepo, catetoryRepo)

	itemID := 1

	itemRepo.On("DeleteItem", itemID).Return(nil)

	err := service.DeleteItem(itemID)

	assert.NoError(t, err)
	itemRepo.AssertExpectations(t)
}

func TestFindItemByID(t *testing.T) {
	itemRepo := new(MockItemRepository)
	catetoryRepo := new(MockCategoryRepository)
	service := NewItemService(itemRepo, catetoryRepo)

	itemID := 1
	item := &Item{ID: itemID, Name: "Item Test"}

	itemRepo.On("FindItemByID", itemID).Return(item, nil)

	result, err := service.FindItemByID(itemID)

	assert.NoError(t, err)
	assert.Equal(t, item, result)
	itemRepo.AssertExpectations(t)
}

func TestListItems(t *testing.T) {
	itemRepo := new(MockItemRepository)
	catetoryRepo := new(MockCategoryRepository)
	service := NewItemService(itemRepo, catetoryRepo)

	items := []*Item{
		{ID: 1, Name: "Item 1"},
		{ID: 2, Name: "Item 2"},
	}

	itemRepo.On("ListItems").Return(items, nil)

	result, err := service.ListItems()

	assert.NoError(t, err)
	assert.Equal(t, items, result)
	itemRepo.AssertExpectations(t)
}
