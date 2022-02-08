package redis

import (
	"encoding/json"

	"github.com/gomodule/redigo/redis"
)

func (p *Pool) StringSet(key string, data interface{}, time int) (bool, error) {
	conn := p.pool.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return false, err
	}
	_, err = conn.Do("SET", key, value)
	if err != nil {
		return false, err
	}
	_, err = p.KeyExpire(key, time)
	if err != nil {
		return false, err
	}
	return true, err
}

func (p *Pool) StringGet(key string) (string, error) {
	conn := p.pool.Get()
	defer conn.Close()

	reply, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}

	return reply, nil
}
func (p *Pool) StringGetInt(key string) (int, error) {
	return redis.Int(p.StringGet(key))
}
