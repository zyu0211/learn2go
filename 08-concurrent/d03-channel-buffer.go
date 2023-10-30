package main

import "fmt"

func main() {

	// 创建一个 可以缓存 2 个值的 通道
    messages := make(chan string, 2)

	// 可以连续 向 通道 中 写入两个值
    messages <- "buffered"
    messages <- "channel"

	// 依次从通道中获取 2 个值
    fmt.Println(<-messages)
    fmt.Println(<-messages)
}