package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCacheRepository struct {
	rdb *redis.Client
}

func NewRedisCacheRepository(url, password string) *RedisCacheRepository {
	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: password,
		DB:       0, // use default DB
	})

	return &RedisCacheRepository{
		rdb: client,
	}
}

func (r *RedisCacheRepository) Get(key string) (string, error) {
	value, err := r.rdb.Get(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", err
	}
	return value, nil
}

func (r *RedisCacheRepository) Set(key string, value string, expiration time.Duration) error {
	err := r.rdb.Set(context.Background(), key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisCacheRepository) Delete(key string) error {
	err := r.rdb.Del(context.Background(), key).Err()
	if err != nil {
		return err
	}
	return nil
}
