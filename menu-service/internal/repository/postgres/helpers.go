package postgres

import (
	"errors"
	"github.com/lib/pq"
)

// QueryParams is used to define the params to StmtNamed queries
type QueryParams map[string]interface{}

// Helper function to check if the error is a UNIQUE constraint violation
func isUniqueViolationError(err error) bool {
	// The PostgreSQL unique key violation error code is 23505
	// It may vary in other database management systems
	// Make sure to check the correct error code for PostgreSQL
	var pgError *pq.Error
	ok := errors.As(err, &pgError)
	if !ok {
		return false
	}
	return pgError.Code == "23505"
}
