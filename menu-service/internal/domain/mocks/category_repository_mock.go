package mocks

import "github.com/rluders/tutorial-microservices/menu-service/internal/domain"

type CategoryRepositoryMock struct {
	data map[int]*domain.Category
}

func NewCategoryRepositoryMock(data map[int]*domain.Category) *CategoryRepositoryMock {
	return &CategoryRepositoryMock{
		data: make(map[int]*domain.Category),
	}
}

func (r *CategoryRepositoryMock) CreateCategory(category *domain.Category) error {
	if category.ID == 0 {
		category.ID = len(r.data) + 1
	}
	r.data[category.ID] = category
	return nil
}

func (r *CategoryRepositoryMock) UpdateCategory(category *domain.Category) error {
	if _, ok := r.data[category.ID]; !ok {
		return domain.ErrCategoryNotFound
	}
	r.data[category.ID] = category
	return nil
}

func (r *CategoryRepositoryMock) DeleteCategory(categoryID int) error {
	if _, ok := r.data[categoryID]; !ok {
		return domain.ErrCategoryNotFound
	}
	delete(r.data, categoryID)
	return nil
}

func (r *CategoryRepositoryMock) FindCategoryByID(categoryID int) (*domain.Category, error) {
	category, ok := r.data[categoryID]
	if !ok {
		return nil, domain.ErrCategoryNotFound
	}
	return category, nil
}

func (r *CategoryRepositoryMock) ListCategories() ([]*domain.Category, error) {
	var categories []*domain.Category
	for _, category := range r.data {
		categories = append(categories, category)
	}
	return categories, nil
}
