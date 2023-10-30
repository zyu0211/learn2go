package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {

	// 使用 os.Setenv 来设置一个键值对
	// 使用 os.Getenv获取一个键对应的值 (如果键不存在，将会返回一个空字符串)
    os.Setenv("FOO", "1")
    fmt.Println("FOO:", os.Getenv("FOO"))
    fmt.Println("BAR:", os.Getenv("BAR"))

    fmt.Println()

	// 使用 os.Environ 来列出所有环境变量键值对
    for _, e := range os.Environ() {
		// 这个函数会返回一个 KEY=value 形式的字符串切片
		// 可以使用 strings.SplitN 来得到键和值
        pair := strings.SplitN(e, "=", 2)
        fmt.Println(pair[0])
    }
}