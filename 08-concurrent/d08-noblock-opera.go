package main

import "fmt"

func main() {
    messages := make(chan string)
    signals := make(chan bool)

	// 非阻塞接收
	// 非缓存通道 messages，如果发送者准备好了，select 会进入 从通道获取 值 的分支
	// 如果 没有准备好的发送者，也不阻塞，而直接到进入 default 分支
    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    default:
        fmt.Println("no message received")
    }

	// 非阻塞发送
	// 非缓存的通道 messages，正常接收者准备好，其发送会阻塞
	// 此处 会进入 default 分支，并不会阻塞在发送分支上
    msg := "hi"
    select {
    case messages <- msg:
        fmt.Println("sent message", msg)
    default:
        fmt.Println("no message sent")
    }

	// 多路的非阻塞的选择器
    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }
}