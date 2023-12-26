package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockCategoryRepository é um mock para CategoryRepository.
type MockCategoryRepository struct {
	mock.Mock
}

func (m *MockCategoryRepository) Create(category *Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *MockCategoryRepository) Update(category *Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *MockCategoryRepository) Delete(categoryID int) error {
	args := m.Called(categoryID)
	return args.Error(0)
}

func (m *MockCategoryRepository) Get(categoryID int) (*Category, error) {
	args := m.Called(categoryID)
	if category := args.Get(0); category != nil {
		return category.(*Category), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockCategoryRepository) List() ([]*Category, error) {
	args := m.Called()
	return args.Get(0).([]*Category), args.Error(1)
}

func (m *MockCategoryRepository) AddItem(itemID, categoryID int) error {
	args := m.Called(itemID, categoryID)
	return args.Error(0)
}

func (m *MockCategoryRepository) RemoveItem(itemID, categoryID int) error {
	args := m.Called(itemID, categoryID)
	return args.Error(0)
}

func (m *MockCategoryRepository) ItemCategories(categoryID int) ([]*Category, error) {
	args := m.Called(categoryID)
	return args.Get(0).([]*Category), args.Error(1)
}

func TestCreateCategory(t *testing.T) {
	repo := new(MockCategoryRepository)
	service := NewCategoryService(repo)

	category := &Category{Name: "Categoria de Teste"}

	repo.On("Create", category).Return(nil)

	err := service.Create(category)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestUpdateCategory(t *testing.T) {
	repo := new(MockCategoryRepository)
	service := NewCategoryService(repo)

	category := &Category{ID: 1, Name: "Categoria de Teste"}

	repo.On("Update", category).Return(nil)

	err := service.Update(category)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestDeleteCategory(t *testing.T) {
	repo := new(MockCategoryRepository)
	service := NewCategoryService(repo)

	categoryID := 1

	repo.On("Delete", categoryID).Return(nil)

	err := service.Delete(categoryID)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestGetCategory(t *testing.T) {
	repo := new(MockCategoryRepository)
	service := NewCategoryService(repo)

	categoryID := 1
	category := &Category{ID: categoryID, Name: "Categoria de Teste"}

	repo.On("Get", categoryID).Return(category, nil)

	result, err := service.Get(categoryID)

	assert.NoError(t, err)
	assert.Equal(t, category, result)
	repo.AssertExpectations(t)
}

func TestListCategories(t *testing.T) {
	repo := new(MockCategoryRepository)
	service := NewCategoryService(repo)

	categories := []*Category{
		{ID: 1, Name: "Categoria 1"},
		{ID: 2, Name: "Categoria 2"},
	}

	repo.On("List").Return(categories, nil)

	result, err := service.List()

	assert.NoError(t, err)
	assert.Equal(t, categories, result)
	repo.AssertExpectations(t)
}
