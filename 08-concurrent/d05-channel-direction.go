package main

import "fmt"

// 当使用通道作为函数的参数时，可以指定这个通道 [只读/只写]

// chan<- : 只写
func ping(pings chan<- string, msg string) {
    pings <- msg
}

// <-chan : 只读
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)

    // 将一个 string 值 写入通道
    ping(pings, "passed message")
    // 从 pings 通道读，并 写入 pongs 通道 
    pong(pings, pongs)
    
    fmt.Println(<-pongs)
}