package bridge

import (
	"net/http"
)

//
// 桥接模式
//    是用于把抽象化与实现化解耦，使得二者可以独立变化。这种类型的设计模式属于结构型模式，
//    它通过提供抽象化和实现化之间的桥接结构，来实现二者的解耦。
//

//请求接口
type Request interface {
	HttpRequest() (*http.Request, error)
}

//cdn请求
type CdnRequest struct {
}

func (cdn *CdnRequest) HttpRequest() (*http.Request, error) {
	return http.NewRequest("GET", "/cdn", nil)
}

//直播请求
type LiveRequest struct {
}

func (cdn *LiveRequest) HttpRequest() (*http.Request, error) {
	return http.NewRequest("GET", "/live", nil)
}

//http客户端
type Client struct {
	Client *http.Client
}

//还是根据传入不同的实例执行不同的方法
func (c *Client) Query(req Request) (resp *http.Response, err error) {
	httpreq, _ := req.HttpRequest()
	resp, err = c.Client.Do(httpreq)
	return
}
