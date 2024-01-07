package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"menu-service/internal/domain"
)

const (
	createCategory    = "create category"
	deleteCategory    = "delete category by id"
	getCategory       = "get category by id"
	listCategory      = "list category"
	updateCategory    = "update category by id"
	addItem           = "add item to category"
	removeItem        = "remove item from category"
	getItemCategories = "get item categories"
)

func queriesCategory() map[string]string {
	return map[string]string{
		createCategory: `INSERT INTO categories (name) VALUES (:name) RETURNING *`,
		deleteCategory: `UPDATE categories SET deleted_at = NOW() WHERE id = :id`,
		getCategory:    `SELECT * FROM categories WHERE id = :id`,
		listCategory:   `SELECT * FROM categories WHERE deleted_at IS NULL ORDER BY name ASC`,
		updateCategory: `UPDATE categories SET name = :name, updated_at = NOW() WHERE id = :id RETURNING *`,
		addItem:        `INSERT INTO item_categories (item_id, category_id) VALUES (:item_id, :category_id)`,
		removeItem:     `DELETE FROM item_categories WHERE item_id = :item_id AND category_id = :category_id`,
		getItemCategories: `SELECT c.id, c.name
		FROM categories c
		INNER JOIN item_categories ic ON c.id = ic.category_id
		WHERE ic.item_id = :item_id`,
	}
}

type CategoryRepository struct {
	*Repository
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	stmts, err := prepareStatements(db, queriesCategory())
	if err != nil {
		log.Fatalf("category repository wasn't able to build all the statements")
		return nil
	}

	return &CategoryRepository{
		&Repository{
			statements: stmts,
			DB:         db,
		},
	}
}

func (r *CategoryRepository) statement(query string) (*sqlx.NamedStmt, error) {
	stmt, ok := r.statements[query]
	if !ok {
		return nil, fmt.Errorf("prepared statement '%s' not found", query)
	}

	return stmt, nil
}

func (r *CategoryRepository) Create(category *domain.Category) error {
	stmt, err := r.statement(createCategory)
	if err != nil {
		return err
	}

	params := QueryParams{
		"name": category.Name,
	}

	if err := stmt.GetContext(context.Background(), category, params); err != nil {
		if isUniqueViolationError(err) {
			return fmt.Errorf("category with name '%s' already exists", category.Name)
		}
		return fmt.Errorf("error creating category: %w", err)
	}

	return nil
}

func (r *CategoryRepository) Update(category *domain.Category) error {
	stmt, err := r.statement(updateCategory)
	if err != nil {
		return err
	}

	params := QueryParams{
		"name": category.Name,
		"id":   category.ID,
	}

	if err := stmt.GetContext(context.Background(), category, params); err != nil {
		if isUniqueViolationError(err) {
			return fmt.Errorf("category with name '%s' already exists", category.Name)
		}
		return fmt.Errorf("error updating category")
	}

	return nil
}

func (r *CategoryRepository) Delete(categoryID int) error {
	stmt, err := r.statement(deleteCategory)
	if err != nil {
		return err
	}

	params := QueryParams{
		"id": categoryID,
	}

	if _, err := stmt.ExecContext(context.Background(), params); err != nil {
		return fmt.Errorf("error deleting category with id '%d'", categoryID)
	}

	return nil
}

func (r *CategoryRepository) Get(categoryID int) (*domain.Category, error) {
	stmt, err := r.statement(getCategory)
	if err != nil {
		return nil, err
	}

	params := QueryParams{
		"id": categoryID,
	}

	category := &domain.Category{}
	if err := stmt.GetContext(context.Background(), category, params); err != nil {
		return nil, fmt.Errorf("error getting the category with id '%d'", categoryID)
	}

	return category, nil
}

func (r *CategoryRepository) List() ([]*domain.Category, error) {
	stmt, err := r.statement(listCategory)
	if err != nil {
		return nil, err
	}

	params := QueryParams{}

	var categories []*domain.Category
	if err := stmt.SelectContext(context.Background(), &categories, params); err != nil {
		return nil, fmt.Errorf("error getting categories")
	}

	return categories, nil
}

func (r *CategoryRepository) AddItem(itemID, categoryID int) error {
	stmt, err := r.statement(addItem)
	if err != nil {
		return err
	}

	params := QueryParams{
		"item_id":     itemID,
		"category_id": categoryID,
	}

	if _, err := stmt.ExecContext(context.Background(), params); err != nil {
		return fmt.Errorf("error adding item '%d' from to category '%d'", itemID, categoryID)
	}

	return nil
}

func (r *CategoryRepository) RemoveItem(itemID, categoryID int) error {
	stmt, err := r.statement(removeItem)
	if err != nil {
		return err
	}

	params := QueryParams{
		"item_id":     itemID,
		"category_id": categoryID,
	}

	if _, err := stmt.ExecContext(context.Background(), params); err != nil {
		return fmt.Errorf("error removing item '%d' from the category '%d'", itemID, categoryID)
	}

	return nil
}
