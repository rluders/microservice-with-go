package domain

import (
	"github.com/pkg/errors"
)

var ErrCategoryNotFound = errors.New("Category not found")

type CategoryService struct {
	categoryRepository CategoryRepository
}

func NewCategoryService(repo CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepository: repo,
	}
}

func (s *CategoryService) CreateCategory(category *Category) error {
	if category == nil {
		return &ValidationError{"Category can't be null"}
	}

	return s.categoryRepository.CreateCategory(category)
}

func (s *CategoryService) UpdateCategory(category *Category) error {
	if category == nil {
		return &ValidationError{"Category can't be null"}
	}

	err := s.categoryRepository.UpdateCategory(category)
	if err != nil {
		return errors.Wrap(err, "Unable to update the category")
	}

	return nil
}

func (s *CategoryService) DeleteCategory(categoryID int) error {
	if categoryID <= 0 {
		return &ValidationError{"Invalid category ID"}
	}

	err := s.categoryRepository.DeleteCategory(categoryID)
	if err != nil {
		if err == ErrCategoryNotFound {
			return ErrCategoryNotFound
		}
		return errors.Wrap(err, "Error to delete the category")
	}

	return nil
}

func (s *CategoryService) FindCategoryByID(categoryID int) (*Category, error) {
	if categoryID <= 0 {
		return nil, &ValidationError{"Invalid category ID"}
	}

	category, err := s.categoryRepository.FindCategoryByID(categoryID)
	if err != nil {
		if err == ErrCategoryNotFound {
			return nil, ErrCategoryNotFound
		}
		return nil, errors.Wrap(err, "Error to find the category")
	}

	return category, nil
}

func (s *CategoryService) ListCategories() ([]*Category, error) {
	categories, err := s.categoryRepository.ListCategories()
	if err != nil {
		return nil, errors.Wrap(err, "Error to list the categories")
	}

	return categories, nil
}
