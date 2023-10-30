package main

import (
	"fmt"
	"time"
)

func main() {

	// 使用 make() 创建一个新通道
	// 此通道 无缓存，存在阻塞
	// 默认发送和接收操作是阻塞的，直到发送方和接收方都就绪
    messages := make(chan string)

	// 在一个 新的 协程 中，向通道中 发送一个 字符串 “ping”
    go func() { 
		time.Sleep(2 * time.Second)	
		messages<- "ping"
	}()

	// 在 main 中，获取 其他协程 发送到 通道 中的值
    msg := <-messages

    fmt.Println(msg)
}