package redis

import (
	"github.com/gomodule/redigo/redis"
)

const (
	INCR = "INCR"
)

type ZsetList []Zset
type Zset struct {
	Key   string
	Score int
}

func (z ZsetList) Len() int { return len(z) }
func (z ZsetList) Less(i, j int) bool {
	return z[i].Score > z[j].Score
}
func (z ZsetList) Swap(i, j int) {
	z[i], z[j] = z[j], z[i]
}

//设置
func ZsetSet(r *redis.Pool, key, value string, score int) error {
	conn := r.Get()
	defer conn.Close()
	_, err := conn.Do("ZADD", key, INCR, score, value)
	if err != nil {
		return err
	}
	return nil
}

//设置 同时自增1
func ZsetSetInccr1(r *redis.Pool, key, value string, score int) error {
	conn := r.Get()
	defer conn.Close()
	_, err := conn.Do("ZADD", key, INCR, 1, value)
	if err != nil {
		return err
	}
	return nil
}

//获取
//从低到高
func ZsetSetGetASC(r *redis.Pool, key string, start, end int) (list ZsetList, err error) {
	conn := r.Get()
	defer conn.Close()
	m, err := redis.IntMap(conn.Do("ZRANGE", key, start, end, "withscores"))
	if err != nil {
		return
	}
	for v := range m {
		list = append(list, Zset{
			Key:   v,
			Score: m[v],
		})
	}
	return
}

//获取
//从高到低
func ZsetSetGetDESC(r *redis.Pool, key string, start, end int) (list ZsetList, err error) {
	conn := r.Get()
	defer conn.Close()
	m, err := redis.IntMap(conn.Do("ZREVRANGE", key, start, end, "withscores"))
	if err != nil {
		return
	}
	for v := range m {
		list = append(list, Zset{
			Key:   v,
			Score: m[v],
		})
	}
	return
}
