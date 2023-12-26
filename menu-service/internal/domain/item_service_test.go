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

func (m *MockItemRepository) Create(item *Item) error {
	args := m.Called(item)
	return args.Error(0)
}

func (m *MockItemRepository) Update(item *Item) error {
	args := m.Called(item)
	return args.Error(0)
}

func (m *MockItemRepository) Delete(itemID int) error {
	args := m.Called(itemID)
	return args.Error(0)
}

func (m *MockItemRepository) Get(itemID int) (*Item, error) {
	args := m.Called(itemID)
	if item := args.Get(0); item != nil {
		return item.(*Item), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockItemRepository) List() ([]*Item, error) {
	args := m.Called()
	return args.Get(0).([]*Item), args.Error(1)
}

func TestCreateItem(t *testing.T) {
	itemRepo := new(MockItemRepository)
	categoryRepo := new(MockCategoryRepository)
	service := NewItemService(itemRepo, categoryRepo)

	item := &Item{Name: "Item Test"}

	itemRepo.On("Create", item).Return(nil)

	err := service.Create(item)

	assert.NoError(t, err)
	itemRepo.AssertExpectations(t)
}

func TestUpdateItem(t *testing.T) {
	itemRepo := new(MockItemRepository)
	categoryRepo := new(MockCategoryRepository)
	service := NewItemService(itemRepo, categoryRepo)

	item := &Item{ID: 1, Name: "Item Test"}

	itemRepo.On("Update", item).Return(nil)

	err := service.Update(item)

	assert.NoError(t, err)
	itemRepo.AssertExpectations(t)
}

func TestDeleteItem(t *testing.T) {
	itemRepo := new(MockItemRepository)
	categoryRepo := new(MockCategoryRepository)
	service := NewItemService(itemRepo, categoryRepo)

	itemID := 1

	itemRepo.On("Delete", itemID).Return(nil)

	err := service.Delete(itemID)

	assert.NoError(t, err)
	itemRepo.AssertExpectations(t)
}

func TestGetItem(t *testing.T) {
	itemRepo := new(MockItemRepository)
	categoryRepo := new(MockCategoryRepository)
	service := NewItemService(itemRepo, categoryRepo)

	itemID := 1
	item := &Item{ID: itemID, Name: "Item Test"}

	itemRepo.On("Get", itemID).Return(item, nil)

	result, err := service.Get(itemID)

	assert.NoError(t, err)
	assert.Equal(t, item, result)
	itemRepo.AssertExpectations(t)
}

func TestListItems(t *testing.T) {
	itemRepo := new(MockItemRepository)
	categoryRepo := new(MockCategoryRepository)
	service := NewItemService(itemRepo, categoryRepo)

	items := []*Item{
		{ID: 1, Name: "Item 1"},
		{ID: 2, Name: "Item 2"},
	}

	itemRepo.On("List").Return(items, nil)

	result, err := service.List()

	assert.NoError(t, err)
	assert.Equal(t, items, result)
	itemRepo.AssertExpectations(t)
}
