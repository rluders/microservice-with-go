package config

import (
	"os"

	"github.com/sherifabdlnaby/configuro"
)

type Config struct {
	Database *Database `validate:"required"`
	Server   struct {
		HTTP *ServerHTTP `validate:"required"`
		GRPC *ServerGRPC `validate:"required"`
	} `validate:"required"`
}

type Database struct {
	Driver   string `validate:"required"`
	Host     string `validate:"required"`
	Port     int    `validate:"required"`
	Username string `validate:"required"`
	Password string `validate:"required"`
	Database string `validate:"required"`
}

type ServerHTTP struct {
	Host     string `validate:"required"`
	Port     int    `validate:"required"`
	UseHTTPS bool
	CertPath string
}

type ServerGRPC struct {
	Port int `validate:"required"`
}

func NewConfig(configPath string) (*Config, error) {
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		return nil, err
	}

	loader, err := configuro.NewConfig(
		configuro.WithLoadFromConfigFile(configPath, false),
		configuro.WithLoadFromEnvVars("APP"))
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = loader.Load(config)
	if err != nil {
		return nil, err
	}

	err = loader.Validate(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
