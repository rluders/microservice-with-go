package domain

import (
	"github.com/pkg/errors"
)

var (
	ErrCategoryNotFound  = errors.New("Category not found")
	ErrCategoryIsNull    = &ValidationError{"Category can't be null"}
	ErrCategoryIDInvalid = &ValidationError{"Invalid category ID"}
)

type CategoryService struct {
	categoryRepository CategoryRepository
}

func NewCategoryService(repo CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepository: repo,
	}
}

func (s *CategoryService) Create(category *Category) error {
	if category == nil {
		return ErrCategoryIsNull
	}

	return s.categoryRepository.Create(category)
}

func (s *CategoryService) Update(category *Category) error {
	if category == nil {
		return ErrCategoryIsNull
	}

	err := s.categoryRepository.Update(category)
	if err != nil {
		return errors.Wrap(err, "Unable to update the category")
	}

	return nil
}

func (s *CategoryService) Delete(categoryID int) error {
	if categoryID <= 0 {
		return ErrCategoryIDInvalid
	}

	err := s.categoryRepository.Delete(categoryID)
	if err != nil {
		if errors.Is(err, ErrCategoryNotFound) {
			return ErrCategoryNotFound
		}
		return errors.Wrap(err, "Error to delete the category")
	}

	return nil
}

func (s *CategoryService) Get(categoryID int) (*Category, error) {
	if categoryID <= 0 {
		return nil, ErrCategoryIDInvalid
	}

	category, err := s.categoryRepository.Get(categoryID)
	if err != nil {
		if errors.Is(err, ErrCategoryNotFound) {
			return nil, ErrCategoryNotFound
		}
		return nil, errors.Wrap(err, "Error to find the category")
	}

	return category, nil
}

func (s *CategoryService) List() ([]*Category, error) {
	categories, err := s.categoryRepository.List()
	if err != nil {
		return nil, errors.Wrap(err, "Error to list the categories")
	}

	return categories, nil
}
