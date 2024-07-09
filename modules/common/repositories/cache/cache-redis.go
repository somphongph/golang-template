package cache

import (
	"golang-template/configs"
	"time"

	"github.com/go-redis/redis"
)

type Cached interface {
	GetCache(string) (string, error)
	SetCache(string, interface{}, int) error
	SetShortCache(string, interface{}) error
	SetLongCache(string, interface{}) error
}

type redisStore struct {
	*redis.Client
	cfg *configs.Redis
}

func InitRedisCache(cfg *configs.Redis) *redisStore {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Host,
		Password: cfg.Pass, // no password set
		DB:       0,        // use default DB
	})

	return &redisStore{rdb, cfg}
}

func (c *redisStore) GetCache(key string) (string, error) {
	k := c.cfg.InstanceName + key
	val, err := c.Get(k).Result()

	return val, err
}

func (c *redisStore) SetCache(key string, value interface{}, duration int) error {
	// Set time in second
	dur := time.Duration(duration) * time.Second
	k := c.cfg.InstanceName + key
	err := c.Set(k, value, dur).Err()

	return err
}

func (c *redisStore) SetShortCache(key string, value interface{}) error {
	dur := time.Duration(c.cfg.ShortCache) * time.Second
	k := c.cfg.InstanceName + key
	err := c.Set(k, value, dur).Err()

	return err
}

func (c *redisStore) SetLongCache(key string, value interface{}) error {
	dur := time.Duration(c.cfg.LongCache) * time.Second
	k := c.cfg.InstanceName + key
	err := c.Set(k, value, dur).Err()

	return err
}
