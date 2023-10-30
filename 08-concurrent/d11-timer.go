package main

import (
    "fmt"
    "time"
)

func main() {

	// 定时器
	// 设置定时器需要等待的时间，会提供一个用于通知的通道
	// 这里的定时器将等待 2 秒
    timer1 := time.NewTimer(2 * time.Second)

	// 定时器通道 会阻塞 2 s，直到 2s 后定时器失效，从其通道获取到通知
    <-timer1.C
    fmt.Println("Timer 1 fired")

    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 fired")
    }()

	// 定时器触发之前将其取消 (与 time.sleep() 的区别)
    isStop := timer2.Stop()
    if isStop {
        fmt.Println("Timer 2 stopped")
    }

    time.Sleep(2 * time.Second)
}