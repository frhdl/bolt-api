package cache

import "github.com/frhdl/bolt-api/app/common"

// RedisConfig cotains the configurations for cache.
type RedisConfig struct {
	Addr     string
	Port     string
	Password string
}

// NewRegisConfig load the configuration for database.
func NewRedisConfig() *RedisConfig {
	password := common.GetEnv("REDIS_PASSWORD", "abcd1234")
	if password == "empty" {
		password = ""
	}

	config := &RedisConfig{
		Addr:     common.GetEnv("REDIS_HOST", "redis:6379"),
		Password: password,
	}

	return config
}
