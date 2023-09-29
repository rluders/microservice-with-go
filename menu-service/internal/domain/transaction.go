package domain

// Transaction defines the interface for database transactions
type Transaction interface {
	Begin() (Transaction, error)
	Commit() error
	Rollback() error
}
