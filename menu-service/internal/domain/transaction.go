// Package domain provides domain-level abstractions and interfaces used by the application.
package domain

// Transaction defines the interface for database transactions.
type Transaction interface {
	// Begin starts a new transaction and returns the corresponding Transaction instance.
	Begin() (Transaction, error)
	// Commit commits the transaction. If an error occurs during the commit, it should be returned.
	Commit() error
	// Rollback rolls back the transaction. If an error occurs during the rollback, it should be returned.
	Rollback() error
}
