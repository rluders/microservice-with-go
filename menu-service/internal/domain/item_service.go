package domain

import (
	"errors"
	"github.com/pkg/errors"
)

var ErrItemNotFound = errors.New("Item not found")

type ItemService struct {
	itemRepository     ItemRepository
	categoryRepository CategoryRepository
	transaction        Transaction
}

func NewItemService(itemRepository ItemRepository, categoryRepository CategoryRepository) *ItemService {
	return &ItemService{
		itemRepository:     itemRepository,
		categoryRepository: categoryRepository,
	}
}

func (s *ItemService) WithTransaction(tx Transaction) *ItemService {
	return &ItemService{
		itemRepository:     s.itemRepository,
		categoryRepository: s.categoryRepository,
		transaction:        tx,
	}
}

func (s *ItemService) CreateItem(item *Item) error {
	if item == nil {
		return &ValidationError{"Item can't be null"}
	}

	if s.transaction != nil {
		tx, err := s.transaction.Begin()
		if err != nil {
			return errors.Wrap(err, "unable to begin transaction")
		}
		defer func() {
			tx.Rollback()
		}()
	}

	if err := s.itemRepository.CreateItem(item); err != nil {
		return err
	}

	for i, c := range item.Categories {
		category, err := s.categoryRepository.FindCategoryByID(c.ID)
		if err != nil {
			return errors.Wrapf(err, "unable to add item to category")
		}

		if err := s.categoryRepository.AddItemToCategory(item.ID, category.ID); err != nil {
			return errors.Wrap(err, "unable to add item to category")
		}
		item.Categories[i] = category
	}

	if s.transaction != nil {
		if err := s.transaction.Commit(); err != nil {
			return errors.Wrap(err, "unable to commit transaction")
		}
	}

	return nil
}

func (s *ItemService) UpdateItem(item *Item) error {
	if item == nil {
		return &ValidationError{"Item can't be null"}
	}

	err := s.itemRepository.UpdateItem(item)
	if err != nil {
		return errors.Wrap(err, "Unable to update the item")
	}

	// TODO: Link Categories

	return nil
}

func (s *ItemService) DeleteItem(itemID int) error {
	if itemID <= 0 {
		return &ValidationError{"Invalid item ID"}
	}

	err := s.itemRepository.DeleteItem(itemID)
	if err != nil {
		if errors.Is(err, ErrItemNotFound) {
			return ErrItemNotFound
		}
		return errors.Wrap(err, "Error to delete the item")
	}

	// TODO: Delete Category Links

	return nil
}

func (s *ItemService) FindItemByID(itemID int) (*Item, error) {
	if itemID <= 0 {
		return nil, &ValidationError{"Invalid item ID"}
	}

	item, err := s.itemRepository.FindItemByID(itemID)
	if err != nil {
		if errors.Is(err, ErrItemNotFound) {
			return nil, ErrItemNotFound
		}
		return nil, errors.Wrap(err, "Error to find the item")
	}

	categories, err := s.categoryRepository.ItemCategories(itemID)
	if err != nil {
		return nil, errors.Wrap(err, "Error to get categories for item")
	}

	item.Categories = categories

	return item, nil
}

func (s *ItemService) ListItems() ([]*Item, error) {
	items, err := s.itemRepository.ListItems()
	if err != nil {
		return nil, errors.Wrap(err, "Error to list items")
	}

	for k, item := range items {
		categories, err := s.categoryRepository.ItemCategories(item.ID)
		if err != nil {
			return nil, errors.Wrap(err, "Error to get categories for item")
		}

		items[k].Categories = categories
	}

	return items, nil
}
