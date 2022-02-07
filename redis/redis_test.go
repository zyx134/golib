package redis

import (
	"fmt"
	"testing"
	"time"
)

func TestSet(t *testing.T) {
	p := New(Config{
		Db:        "2",
		Port:      9999,
		MaxIdle:   10,
		MaxActive: 20,
		Timeout:   time.Duration(time.Second * 30),
	})
	// p.SetSet("test:test", "ttt")
	p.SetSet("test:test", "tt")
	p.SetSet("test:test", "ttt")
	p.SetSet("test:test", "ttt2")
	p.SetSet("test:test", "ttt3")
	p.SetSet("test:test", "ttt4")
	p.SetSet("test:test", "ttt5")
	fmt.Println("是否存在", p.SetExist("test:test", "tt"))
	fmt.Println("del", p.SetDel("test:test", "tt"))
	fmt.Println("all", p.SetGetAll("test:test"))

}
