package redis

import (
	"encoding/json"

	"github.com/gomodule/redigo/redis"
	"github.com/zyx134/golib/_string"
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
	conn.Do("EXPIRE", key, time)

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
func (p *Pool) StringSetInt(key string, data int, time int) (bool, error) {
	conn := p.pool.Get()
	defer conn.Close()
	reply, err := redis.Bool(conn.Do("SET", key, data))
	conn.Do("EXPIRE", key, time)

	return reply, err
}
func (p *Pool) StringGetInt(key string) (ret int, err error) {
	s, err := p.StringGet(key)
	if err != nil {
		return
	}
	return _string.ToInt(s)
}
