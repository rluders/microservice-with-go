package domain

import (
	"time"
)

// Item represents an item entity with its associated properties.
type Item struct {
	ID          int         `json:"id" db:"id"`
	Name        string      `json:"name" db:"name"`
	Description string      `json:"description" db:"description"`
	Price       float64     `json:"price" db:"price"`
	Categories  []*Category `json:"categories,omitempty" db:"categories"`
	CreatedAt   time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time  `json:"deleted_at,omitempty" db:"deleted_at"`
}
