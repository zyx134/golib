package bridge

import (
	"fmt"
	"net/http"
	"testing"
)

func TestBridge(t *testing.T) {

	//这个客户端可以理解成就是桥
	client := &Client{Client: http.DefaultClient}

	//提供抽象化和实现化之间的桥接结构
	//传入不同的实例即可实现不同的操作

	//这里传入cdn实例，执行的是cdn的HttpRequest方法
	cdnReq := &CdnRequest{}
	fmt.Println(client.Query(cdnReq))

	//传入live实例，执行的是live的HttpRequest方法
	liveReq := &LiveRequest{}
	fmt.Println(client.Query(liveReq))
}
