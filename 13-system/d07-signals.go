package main

import (
    "fmt"
    "os"
	// Go 中使用通道来处理信号
    "os/signal"
    "syscall"
)

func main() {

	// 通过向一个通道发送 os.Signal 值来发送信号通知
	// 注意: 这个通道应该被缓存
    sigs := make(chan os.Signal, 1)

	// signal.Notify 注册给定的通道，用于接收特定信号
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    done := make(chan bool, 1)

	// 这个协程执行一个阻塞的信号接收操作
	// 当它接收到一个值时，它将打印这个值，然后通知程序可以退出
    go func() {

        sig := <-sigs
        fmt.Println()
        fmt.Println(sig)
        done <- true
    }()

    fmt.Println("awaiting signal")
	// 程序将在这里进行等待，直到它得到了期望的信号
    <-done
    fmt.Println("exiting")
}