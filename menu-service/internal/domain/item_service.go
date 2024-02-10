package domain

import (
	"github.com/pkg/errors"
	"log"
)

var (
	ErrItemNotFound  = errors.New("Item not found")
	ErrItemIsNull    = &ValidationError{"Item can't be null"}
	ErrItemIDInvalid = &ValidationError{"Invalid item ID"}
)

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

func (s *ItemService) Create(item *Item) error {
	if item == nil {
		return ErrCategoryIsNull
	}

	if s.transaction != nil {
		tx, err := s.transaction.Begin()
		if err != nil {
			return errors.Wrap(err, "Unable to begin transaction")
		}
		defer func() {
			if err = tx.Rollback(); err != nil {
				log.Printf("Fail to rollback item creation transaction: %v", err)
			}
		}()
	}

	if err := s.itemRepository.Create(item); err != nil {
		return err
	}

	//for i, c := range item.Categories {
	//	category, err := s.categoryRepository.Get(c.ID)
	//	if err != nil {
	//		return errors.Wrapf(err, "Unable to add item to category")
	//	}
	//
	//	if err := s.categoryRepository.AddItem(item.ID, category.ID); err != nil {
	//		return errors.Wrap(err, "Unable to add item to category")
	//	}
	//	item.Categories[i] = category
	//}

	if s.transaction != nil {
		if err := s.transaction.Commit(); err != nil {
			return errors.Wrap(err, "Unable to commit transaction")
		}
	}

	return nil
}

func (s *ItemService) Update(item *Item) error {
	if item == nil {
		return ErrItemIsNull
	}

	err := s.itemRepository.Update(item)
	if err != nil {
		return errors.Wrap(err, "Unable to update the item")
	}

	return nil
}

func (s *ItemService) Delete(itemID int) error {
	if itemID <= 0 {
		return ErrItemIDInvalid
	}

	err := s.itemRepository.Delete(itemID)
	if err != nil {
		if errors.Is(err, ErrItemNotFound) {
			return ErrItemNotFound
		}
		return errors.Wrap(err, "Error to delete the item")
	}

	return nil
}

func (s *ItemService) Get(itemID int) (*Item, error) {
	if itemID <= 0 {
		return nil, ErrItemIDInvalid
	}

	item, err := s.itemRepository.Get(itemID)
	if err != nil {
		if errors.Is(err, ErrItemNotFound) {
			return nil, ErrItemNotFound
		}
		return nil, errors.Wrap(err, "Error to find the item")
	}

	return item, nil
}

func (s *ItemService) List() ([]*Item, error) {
	items, err := s.itemRepository.List()
	if err != nil {
		return nil, errors.Wrap(err, "Error to list items")
	}

	return items, nil
}
