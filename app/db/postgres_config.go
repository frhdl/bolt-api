package db

import "github.com/frhdl/bolt-api/app/common"

// PostgresConfig contains the configurations for database.
type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SSL      string
}

// NewPostgresConfig load the configuration for database.
func NewPostgresConfig() *PostgresConfig {
	config := &PostgresConfig{
		Host:     common.GetEnv("DATABASE_HOST", "postgres"),
		Port:     common.GetEnv("DATABASE_PORT", "5432"),
		User:     common.GetEnv("DATABASE_USER", "postgres"),
		Password: common.GetEnv("DATABASE_PASSWORD", ""),
		DbName:   common.GetEnv("DATABASE_DBNAME", "postgres"),
		SSL:      common.GetEnv("SSL", "disable"),
	}

	return config
}
