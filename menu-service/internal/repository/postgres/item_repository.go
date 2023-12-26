package postgres

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"menu-service/internal/domain"
)

const (
	createItem = "create item"
	deleteItem = "delete item by id"
	getItem    = "get item by id"
	listItem   = "list item"
	updateItem = "update item by id"
)

func queriesItem() map[string]string {
	return map[string]string{
		createItem: `INSERT INTO items (name, description, price) VALUES ($1, $2, $3) RETURNING *`,
		deleteItem: `UPDATE items SET deleted_at = NOW() WHERE id = $1`,
		getItem:    `SELECT * FROM items WHERE id = $1`,
		listItem:   `SELECT * FROM items WHERE deleted_at IS NULL ORDER BY name ASC`,
		updateItem: `UPDATE items SET name = $1, description = $2, price = $3, updated_at = NOW() WHERE id = $4 RETURNING *`,
	}
}

type ItemRepository struct {
	DB         *sqlx.DB
	statements map[string]*sqlx.Stmt
}

func NewItemRepository(db *sqlx.DB) *ItemRepository {
	sqlStatements := make(map[string]*sqlx.Stmt)

	var errs []error
	for queryName, query := range queriesItem() {
		stmt, err := db.Preparex(query)
		if err != nil {
			log.Printf("error preparing statement %s: %v", queryName, err)
			errs = append(errs, err)
		}
		sqlStatements[queryName] = stmt
	}

	if len(errs) > 0 {
		log.Fatalf("item repository wasn't able to build all the statements")
		return nil
	}

	return &ItemRepository{
		DB:         db,
		statements: sqlStatements,
	}
}

func (r *ItemRepository) statement(query string) (*sqlx.Stmt, error) {
	stmt, ok := r.statements[query]
	if !ok {
		return nil, fmt.Errorf("prepared statement '%s' not found", query)
	}

	return stmt, nil
}

func (r *ItemRepository) Create(item *domain.Item) error {
	stmt, err := r.statement(createItem)
	if err != nil {
		return err
	}

	if err := stmt.Get(item, item.Name, item.Description, item.Price); err != nil {
		return fmt.Errorf("error creating item: %v", err)
	}

	return nil
}

func (r *ItemRepository) Update(item *domain.Item) error {
	stmt, err := r.statement(updateItem)
	if err != nil {
		return err
	}

	item.UpdatedAt = time.Now()

	params := []interface{}{
		item.Name,
		item.Description,
		item.Price,
		item.ID,
	}

	if err := stmt.Get(item, params...); err != nil {
		return fmt.Errorf("error updating item")
	}

	return nil
}

func (r *ItemRepository) Delete(itemID int) error {
	stmt, err := r.statement(deleteItem)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(itemID); err != nil {
		return fmt.Errorf("error deleting item with id '%d'", itemID)
	}

	return nil
}

func (r *ItemRepository) Get(itemID int) (*domain.Item, error) {
	stmt, err := r.statement(getItem)
	if err != nil {
		return nil, err
	}

	item := &domain.Item{}
	if err := stmt.Get(item, itemID); err != nil {
		return nil, fmt.Errorf("error getting the item with id '%d'", itemID)
	}

	return item, nil
}

func (r *ItemRepository) List() ([]*domain.Item, error) {
	stmt, err := r.statement(listItem)
	if err != nil {
		return nil, err
	}

	var items []*domain.Item
	if err := stmt.Select(&items); err != nil {
		return nil, fmt.Errorf("error getting items")
	}

	return items, nil
}
