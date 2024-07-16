package main

import (
	"fmt"
	"net/http"
)

// handlers 是 net/http 服务器里面的一个基本概念。
// handler 对象实现了 http.Handler 接口。
// 编写 handler 的常见方法是，在具有适当签名的函数上使用 http.HandlerFunc 适配器

// handler 函数有两个参数，http.ResponseWriter 和 http.Request。
func helloworld(w http.ResponseWriter, req *http.Request) {

	// response writer 被用于写入 HTTP 响应数据
	fmt.Fprintf(w, "hello\n")
}

// 读取的 HTTP 请求 header 中的所有内容，并将他们输出至 response body
func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	// http.HandleFunc 函数，将我们的 handler 注册到服务器路由
	http.HandleFunc("/hello", helloworld)
	http.HandleFunc("/headers", headers)

	// 调用 ListenAndServe 并带上端口和 handler
	// nil 表示使用我们刚刚设置的默认路由器
	http.ListenAndServe(":8090", nil)
}
