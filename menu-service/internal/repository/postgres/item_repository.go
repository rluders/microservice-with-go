package postgres

import (
	"context"
	"fmt"
	"github.com/jackskj/carta"
	"github.com/jmoiron/sqlx"
	"log"
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
		createItem: `INSERT INTO 
    		items (
    		       name, 
    		       description, 
    		       price
			) 
			VALUES (
			        :name, 
			        :description, 
			        :price
			) 
			RETURNING *`,
		deleteItem: `UPDATE 
    		items 
		SET 
			deleted_at = NOW() 
		WHERE 
			id = :id`,
		getItem: `SELECT
            i.id AS id,
            i.name AS name,
            i.description AS description,
            i.price AS price,
            i.created_at AS created_at,
            i.updated_at AS updated_at,
			i.deleted_at AS deleted_at,
			c.id AS "categories_id", 
			c.name AS "categories_name", 
			c.created_at AS "categories_created_at", 
			c.updated_at AS "categories_updated_at", 
			c.deleted_at AS "categories_deleted_at"
        FROM
            items i
        LEFT JOIN
            item_categories ic ON i.id = ic.item_id
        LEFT JOIN
            categories c ON ic.category_id = c.id
        WHERE
            i.id = :id
		AND i.deleted_at IS NULL
		ORDER BY c.name`,
		listItem: `SELECT
					i.name,
					i.id,
					i.description,
					i.price,
					i.created_at,
					i.updated_at,
					c.id AS "categories_id", 
					c.name AS "categories_name", 
					c.created_at AS "categories_created_at", 
					c.updated_at AS "categories_updated_at", 
					c.deleted_at AS "categories_deleted_at"
				FROM
					items i
				LEFT JOIN
					item_categories ic ON i.id = ic.item_id
				LEFT JOIN
					categories c ON ic.category_id = c.id
				WHERE
					i.deleted_at IS NULL
				ORDER BY
					i.name`,
		updateItem: `UPDATE 
    		items 
		SET 
			name = :name, 
			description = :description, 
			price = :price, 
			updated_at = NOW() 
		 WHERE 
		     id = :id
		 RETURNING *`,
	}
}

type ItemRepository struct {
	DB         *sqlx.DB
	statements map[string]*sqlx.NamedStmt
}

func NewItemRepository(db *sqlx.DB) *ItemRepository {
	sqlStatements := make(map[string]*sqlx.NamedStmt)

	var errs []error
	for queryName, query := range queriesItem() {
		stmt, err := db.PrepareNamed(query)
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

func (r *ItemRepository) statement(query string) (*sqlx.NamedStmt, error) {
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

	params := &QueryParams{
		"name":        item.Name,
		"description": item.Description,
		"price":       item.Price,
	}

	if err := stmt.GetContext(context.Background(), item, params); err != nil {
		return fmt.Errorf("error creating item: %v", err)
	}

	return nil
}

func (r *ItemRepository) Update(item *domain.Item) error {
	stmt, err := r.statement(updateItem)
	if err != nil {
		return err
	}

	params := &QueryParams{
		"item":        item.Name,
		"description": item.Description,
		"price":       item.Price,
		"id":          item.ID,
	}
	if err := stmt.GetContext(context.Background(), item, params); err != nil {
		return fmt.Errorf("error updating item")
	}

	return nil
}

func (r *ItemRepository) Delete(itemID int) error {
	stmt, err := r.statement(deleteItem)
	if err != nil {
		return err
	}

	params := &QueryParams{
		"id": itemID,
	}
	if _, err := stmt.ExecContext(context.Background(), params); err != nil {
		return fmt.Errorf("error deleting item with id '%d'", itemID)
	}

	return nil
}

func (r *ItemRepository) Get(itemID int) (*domain.Item, error) {
	stmt, err := r.statement(getItem)
	if err != nil {
		return nil, err
	}

	params := QueryParams{
		"id": itemID,
	}

	rows, err := stmt.QueryxContext(context.Background(), params)
	if err != nil {
		return nil, fmt.Errorf("error getting the item with id '%d'", itemID)
	}
	defer func(rows *sqlx.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf("error to close rows: %v", err)
		}
	}(rows)

	item := &domain.Item{}
	err = carta.Map(rows.Rows, item)
	if err != nil {
		return nil, fmt.Errorf("unable to map result: %v", err)
	}

	return item, nil
}

func (r *ItemRepository) List() ([]*domain.Item, error) {
	stmt, err := r.statement(listItem)
	if err != nil {
		return nil, err
	}

	params := QueryParams{}

	rows, err := stmt.QueryxContext(context.Background(), params)
	if err != nil {
		return nil, fmt.Errorf("error getting the item list")
	}
	defer func(rows *sqlx.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf("error to close rows: %v", err)
		}
	}(rows)

	var items []*domain.Item
	err = carta.Map(rows.Rows, &items)
	if err != nil {
		return nil, fmt.Errorf("unable to map results: %v", err)
	}

	return items, nil
}
