package ports

import (
	"database/sql"
	"time"

	"github.com/go-redis/redis/v8"
)

// Persistence interface.
type Persistence interface {
	Exec(string, ...interface{}) (bool, error)
	QueryRow(string, ...interface{}) (*sql.Rows, error)
	Shutdown()
	RunFileQuery(string) error
}

// Cache interface.
type Cache interface {
	SetValue(string, interface{}, time.Duration) error
	GetValue(string) (string, error)
	GetInstance() *redis.Client
	ExpireKey(string) error
}
