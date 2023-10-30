package main

import (
    "fmt"
    "time"
)

func main() {

	// 打点器: 以固定的时间间隔重复执行
	// 新建 一个打点器，每 500 ms 发送一次任务
	// 也是 通过 通道的形式 实现的
    ticker := time.NewTicker(500 * time.Millisecond)
    done := make(chan bool)

    go func() {
		// 死循环，不断使用 select 进行选择
        for {
            select {
            case <-done:
                return
			// 每 500 ms，收到 打点器 的通知
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
            }
        }
    }()

    time.Sleep(1600 * time.Millisecond)
	// 手动 停止 打点器
	// 打点器一旦停止，将不能再从它的通道中接收到值
    ticker.Stop()
    done <- true
    fmt.Println("Ticker stopped")
}