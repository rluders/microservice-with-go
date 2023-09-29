package postgres

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rluders/tutorial-microservices/menu-service/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConnectWithMock(t *testing.T) {
	cfg := &config.Database{
		Host:     "localhost",
		Port:     5432,
		Database: "test_db",
		Username: "test_user",
		Password: "test_password",
		Driver:   "postgres",
	}

	// Crie um objeto sql.DB simulado e um mock para a conexão
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	mock.ExpectPing()

	// Chame a função Connect com o DBWrapper
	_, err = Connect(cfg)
	require.NoError(t, err)

	// Verifique se todas as expectativas foram atendidas
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
