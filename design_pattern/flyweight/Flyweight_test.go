package flyweight

import "testing"

func TestFlyweight(t *testing.T) {
	pool := NewDbConnectPool(2)
	conn := pool.Get()
	conn.Do()
	pool.Put(conn)
}
