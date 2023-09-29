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

func (m *MockCategoryRepository) CreateCategory(category *Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *MockCategoryRepository) UpdateCategory(category *Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *MockCategoryRepository) DeleteCategory(categoryID int) error {
	args := m.Called(categoryID)
	return args.Error(0)
}

func (m *MockCategoryRepository) FindCategoryByID(categoryID int) (*Category, error) {
	args := m.Called(categoryID)
	if category := args.Get(0); category != nil {
		return category.(*Category), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockCategoryRepository) ListCategories() ([]*Category, error) {
	args := m.Called()
	return args.Get(0).([]*Category), args.Error(1)
}

func TestCreateCategory(t *testing.T) {
	repo := new(MockCategoryRepository)
	service := NewCategoryService(repo)

	category := &Category{Name: "Categoria de Teste"}

	repo.On("CreateCategory", category).Return(nil)

	err := service.CreateCategory(category)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestUpdateCategory(t *testing.T) {
	repo := new(MockCategoryRepository)
	service := NewCategoryService(repo)

	category := &Category{ID: 1, Name: "Categoria de Teste"}

	repo.On("UpdateCategory", category).Return(nil)

	err := service.UpdateCategory(category)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestDeleteCategory(t *testing.T) {
	repo := new(MockCategoryRepository)
	service := NewCategoryService(repo)

	categoryID := 1

	repo.On("DeleteCategory", categoryID).Return(nil)

	err := service.DeleteCategory(categoryID)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestFindCategoryByID(t *testing.T) {
	repo := new(MockCategoryRepository)
	service := NewCategoryService(repo)

	categoryID := 1
	category := &Category{ID: categoryID, Name: "Categoria de Teste"}

	repo.On("FindCategoryByID", categoryID).Return(category, nil)

	result, err := service.FindCategoryByID(categoryID)

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

	repo.On("ListCategories").Return(categories, nil)

	result, err := service.ListCategories()

	assert.NoError(t, err)
	assert.Equal(t, categories, result)
	repo.AssertExpectations(t)
}
