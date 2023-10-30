package main

import (
    "fmt"
    "net"
    "net/url"
)

func main() {

    s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// 解析这个 URL 并确保解析没有出错
    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }

    fmt.Println(u.Scheme)

    fmt.Println(u.User)
    fmt.Println(u.User.Username())
    p, _ := u.User.Password()
    fmt.Println(p)

	// Host 包含主机名和端口号
    fmt.Println(u.Host)
    host, port, _ := net.SplitHostPort(u.Host)
    fmt.Println(host)
    fmt.Println(port)

	// 提取路径和 # 后面的查询片段信息
    fmt.Println(u.Path)
    fmt.Println(u.Fragment)

	// 要得到字符串中的 k=v 这种格式的查询参数，可以使用 RawQuery 函数
    fmt.Println(u.RawQuery)
	// 将查询参数解析为一个 map
    m, _ := url.ParseQuery(u.RawQuery)
    fmt.Println(m)
    fmt.Println(m["k"][0])
}