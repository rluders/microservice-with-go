package postgres

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rluders/tutorial-microservices/menu-service/internal/domain"
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
		createCategory: `INSERT INTO categories (name) VALUES ($1) RETURNING *`,
		deleteCategory: `UPDATE categories SET deleted_at = NOW() WHERE id = $1`,
		getCategory:    `SELECT * FROM categories WHERE id = $1`,
		listCategory:   `SELECT * FROM categories WHERE deleted_at IS NULL ORDER BY name ASC`,
		updateCategory: `UPDATE categories SET name = $1, updated_at = NOW() WHERE id = $2 RETURNING *`,
		addItem:        `INSERT INTO item_categories (item_id, category_id) VALUES ($1, $2)`,
		removeItem:     `DELETE FROM item_categories WHERE item_id = $1 AND category_id = $2`,
		getItemCategories: `SELECT c.id, c.name
		FROM categories c
		INNER JOIN item_categories ic ON c.id = ic.category_id
		WHERE ic.item_id = $1`,
	}
}

type CategoryRepository struct {
	DB         *sqlx.DB
	statements map[string]*sqlx.Stmt
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	sqlStatements := make(map[string]*sqlx.Stmt)

	var errs []error
	for queryName, query := range queriesCategory() {
		stmt, err := db.Preparex(query)
		if err != nil {
			log.Printf("error preparing statement %s: %v", queryName, err)
			errs = append(errs, err)
		}
		sqlStatements[queryName] = stmt
	}

	if len(errs) > 0 {
		log.Fatalf("category repository wasn't able to build all the statements")
		return nil
	}

	return &CategoryRepository{
		DB:         db,
		statements: sqlStatements,
	}
}

func (r *CategoryRepository) CreateCategory(category *domain.Category) error {
	stmt, ok := r.statements[createCategory]
	if !ok {
		return fmt.Errorf("prepared statement '%s' not found", createCategory)
	}

	if err := stmt.Get(category, category.Name); err != nil {
		if isUniqueViolationError(err) {
			return fmt.Errorf("category with name '%s' already exists", category.Name)
		}
		return fmt.Errorf("error creating category: %v", err)
	}

	return nil
}

func (r *CategoryRepository) UpdateCategory(category *domain.Category) error {
	stmt, ok := r.statements[updateCategory]
	if !ok {
		return fmt.Errorf("prepared statement '%s' not found", updateCategory)
	}

	category.UpdatedAt = time.Now()

	params := []interface{}{
		category.Name,
		category.ID,
	}

	if err := stmt.Get(category, params...); err != nil {
		if isUniqueViolationError(err) {
			return fmt.Errorf("category with name '%s' already exists", category.Name)
		}
		return fmt.Errorf("error updating category")
	}

	return nil
}

func (r *CategoryRepository) DeleteCategory(categoryID int) error {
	stmt, ok := r.statements[deleteCategory]
	if !ok {
		return fmt.Errorf("prepared statement '%s' not found", deleteCategory)
	}

	if _, err := stmt.Exec(categoryID); err != nil {
		return fmt.Errorf("error deleting category with id '%d'", categoryID)
	}

	return nil
}

func (r *CategoryRepository) FindCategoryByID(categoryID int) (*domain.Category, error) {
	stmt, ok := r.statements[getCategory]
	if !ok {
		return nil, fmt.Errorf("prepared statement '%s' not found", getCategory)
	}

	category := &domain.Category{}
	if err := stmt.Get(category, categoryID); err != nil {
		return nil, fmt.Errorf("error getting the category with id '%d'", categoryID)
	}

	return category, nil
}

func (r *CategoryRepository) ListCategories() ([]*domain.Category, error) {
	stmt, ok := r.statements[listCategory]
	if !ok {
		return nil, fmt.Errorf("prepared statement '%s' not found", listCategory)
	}

	var categories []*domain.Category
	if err := stmt.Select(&categories); err != nil {
		return nil, fmt.Errorf("error getting categories")
	}

	return categories, nil
}

func (r *CategoryRepository) AddItemToCategory(itemID, categoryID int) error {
	stmt, ok := r.statements[addItem]
	if !ok {
		return fmt.Errorf("prepared statement '%s' not found", addItem)
	}

	if _, err := stmt.Exec(itemID, categoryID); err != nil {
		return fmt.Errorf("error adding item '%d' from to category '%d'", itemID, categoryID)
	}

	return nil
}

func (r *CategoryRepository) RemoveItemFromCategory(itemID, categoryID int) error {
	stmt, ok := r.statements[removeItem]
	if !ok {
		return fmt.Errorf("prepared statement '%s' not found", removeItem)
	}

	if _, err := stmt.Exec(itemID, categoryID); err != nil {
		return fmt.Errorf("error removing item '%d' from the category '%d'", itemID, categoryID)
	}

	return nil
}

func (r *CategoryRepository) ItemCategories(itemID int) ([]*domain.Category, error) {
	stmt, ok := r.statements[getItemCategories]
	if !ok {
		return nil, fmt.Errorf("prepared statement '%s' not found", getItemCategories)
	}

	var categories []*domain.Category

	if err := stmt.Select(&categories, itemID); err != nil {
		return nil, err
	}

	return categories, nil
}
