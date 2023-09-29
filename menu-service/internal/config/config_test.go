package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewConfigValidFile(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "config.*.yaml")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	configContent := `
database:
  driver: postgres
  host: localhost
  port: 5432
  username: test_user
  password: test_password
  database: test_db
server:
  http:
    host: localhost
    port: 8080
  grpc:
    port: 50051
`
	fmt.Println(tmpFile.Name())
	err = os.WriteFile(tmpFile.Name(), []byte(configContent), 0644)
	require.NoError(t, err)

	config, err := NewConfig(tmpFile.Name())
	require.NoError(t, err)

	assert.Equal(t, "postgres", config.Database.Driver)
	assert.Equal(t, "localhost", config.Database.Host)
	assert.Equal(t, 5432, config.Database.Port)
	assert.Equal(t, "test_user", config.Database.Username)
	assert.Equal(t, "test_password", config.Database.Password)
	assert.Equal(t, "test_db", config.Database.Database)
	assert.Equal(t, "localhost", config.Server.HTTP.Host)
	assert.Equal(t, 8080, config.Server.HTTP.Port)
	assert.Equal(t, 50051, config.Server.GRPC.Port)
}

func TestNewConfigInvalidFile(t *testing.T) {
	config, err := NewConfig("nonexistent.yaml")
	assert.Error(t, err)
	assert.Nil(t, config)
}
