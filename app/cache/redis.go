package cache

import (
	systemcontext "context"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq" // postgres
)

// Redis contains objects for database communication.
type Redis struct {
	*redis.Client
}

// NewRedis create new cache instance.
func NewRedis(config *RedisConfig) (*Redis, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password, // no password set
		DB:       1,               // use default DB
	})

	systemctx := systemcontext.Background()
	err := rdb.Ping(systemctx).Err()
	if err != nil {
		return nil, err
	}

	return &Redis{rdb}, nil
}

//SetValue set expiration time of key value.
func (cache *Redis) SetValue(key string, value interface{}, expiration time.Duration) error {
	systemctx := systemcontext.Background()
	err := cache.Set(systemctx, key, value, expiration).Err()
	return err
}

//GetValue return a key value.
func (cache *Redis) GetValue(key string) (string, error) {
	systemctx := systemcontext.Background()
	val, err := cache.Get(systemctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

//ExpireKey expire a key.
func (cache *Redis) ExpireKey(key string) error {
	systemctx := systemcontext.Background()
	err := cache.Expire(systemctx, key, 1*time.Second).Err()
	if err != nil {
		return err
	}

	return nil
}

//GetInstance return a new cache client.
func (cache *Redis) GetInstance() *redis.Client {
	return cache.Client
}
