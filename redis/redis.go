package redis

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/zyx134/golib/_string"

	"github.com/gomodule/redigo/redis"
)

func New(host, password, db string, port, maxIdle, maxActive int, idleTimeout time.Duration) *redis.Pool {
	r := new(redis.Pool)
	r.MaxIdle = maxIdle
	r.MaxActive = maxActive
	r.IdleTimeout = idleTimeout
	r.Dial = func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
		if err != nil {
			return nil, err
		}
		if password != "" {
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
		}
		_, err = c.Do("SELECT", db)
		if err != nil {
			c.Close()
			return nil, err
		}
		return c, err
	}
	r.TestOnBorrow = func(c redis.Conn, t time.Time) error {
		_, err := c.Do("PING")
		return err
	}

	return r
}
func Exists(r *redis.Pool, key string) bool {
	conn := r.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

func Delete(r *redis.Pool, key string) (bool, error) {
	conn := r.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

func LikeDeletes(r *redis.Pool, key string) error {
	conn := r.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(r, key)
		if err != nil {
			return err
		}
	}

	return nil
}
func Set(r *redis.Pool, key string, data interface{}, time int) (bool, error) {
	conn := r.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return false, err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return false, err
	}
	conn.Do("EXPIRE", key, time)

	return true, err
}

func Get(r *redis.Pool, key string) (string, error) {
	conn := r.Get()
	defer conn.Close()

	reply, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}

	return reply, nil
}
func SetInt(r *redis.Pool, key string, data int, time int) (bool, error) {
	conn := r.Get()
	defer conn.Close()
	reply, err := redis.Bool(conn.Do("SET", key, data))
	conn.Do("EXPIRE", key, time)

	return reply, err
}
func GetInt(r *redis.Pool, key string) (ret int, err error) {
	s, err := Get(r, key)
	if err != nil {
		return
	}
	return _string.ToInt(s)
}
