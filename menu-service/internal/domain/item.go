package domain

import "time"

type Item struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Price       float64     `json:"price"`
	Categories  []*Category `json:"categories"`
	CreatedAt   time.Time   `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time   `db:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time  `db:"deleted_at" json:"deleted_at,omitempty"`
}
