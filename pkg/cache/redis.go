package cache

import (
	"encoding/json"
	"time"

	"github.com/gomodule/redigo/redis"
)

type RedisClient struct {
	pool *redis.Pool
}

func NewRedisClient(host string, password string, maxIdle, maxActive int, idleTimeout time.Duration) *RedisClient {
	pool := &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return &RedisClient{pool: pool}
}

func (r *RedisClient) Set(key string, data interface{}, expiration int) error {
	conn := r.pool.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		// log.Printf("[RedisClient] failed to marshal data: %v", err)
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		// log.Printf("[RedisClient] failed to set key: %s, error: %v", key, err)
		return err
	}

	if expiration > 0 {
		_, err = conn.Do("EXPIRE", key, expiration)
		if err != nil {
			// log.Printf("[RedisClient] failed to set expiration for key: %s, error: %v", key, err)
			return nil
		}
	}

	// log.Printf("[RedisClient] successfully set key: %s, expiration: %d", key, expiration)
	return err
}

func (r *RedisClient) Get(key string) ([]byte, error) {
	conn := r.pool.Get()
	defer conn.Close()

	return redis.Bytes(conn.Do("GET", key))
}

func (r *RedisClient) Exists(key string) bool {
	conn := r.pool.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	return err == nil && exists
}

func (r *RedisClient) Delete(key string) (bool, error) {
	conn := r.pool.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}
