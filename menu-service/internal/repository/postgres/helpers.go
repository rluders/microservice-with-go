package postgres

import (
	"errors"
	"github.com/lib/pq"
)

// Função auxiliar para verificar se o erro é uma violação da restrição UNIQUE
func isUniqueViolationError(err error) bool {
	// O código de erro de violação de chave única do PostgreSQL é 23505
	// Pode variar em outros sistemas de gerenciamento de banco de dados
	// Certifique-se de verificar o código de erro correto para o PostgreSQL
	var pgError *pq.Error
	ok := errors.As(err, &pgError)
	if !ok {
		return false
	}
	return pgError.Code == "23505"
}
