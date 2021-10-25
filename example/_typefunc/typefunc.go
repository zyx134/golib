package _typefunc

import "net/http"

func Http() {
	http.HandleFunc("/", h)
	http.ListenAndServe("0.0.0.0:1111", nil)
}
func h(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello word"))
}

//分析

/**
h函数的函数签名为 func(ResponseWriter, *Request)
和http.HandlerFunc签名一致
1 在net/http/server.go 的2512行
DefaultServeMux 多路复用器
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}
2 在net/http/server.go 的2497行
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	if handler == nil {
		panic("http: nil handler")
	}
	mux.Handle(pattern, HandlerFunc(handler))
}
3 在net/http/server.go 的2453行
路由注册
等待函数调用

HandlerFunc(handler)
将h函数转换为HandlerFunc
这样h函数就实现了Handler接口
type HandlerFunc func(ResponseWriter, *Request)
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}
h函数就可以调用ServeHTTP函数
调用ServeHTTP函数的结果就是调用h函数自己
这样用户就可以不用手动写ServeHTTP函数
*/
