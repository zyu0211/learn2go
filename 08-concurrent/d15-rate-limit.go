package main

import (
    "fmt"
    "time"
)

// Go 可以 基于协程、通道、打点器，实现速率限制
func main() {

	// 声名一个通道 requests: 存放请求
    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)

	// limiter 通道每 200ms 接收一个值
	// 作为 任务速率限制 的调度器
    limiter := time.Tick(200 * time.Millisecond)

	// 遍历 通道，获取任务
    for req := range requests {
		// 每次 处理任务前，需要从 limiter 通道获取 一个值， 
		// limiter 每 200ms 会生成一个值(不然会阻塞)，因此 任务 200ms 执行一次
        <-limiter
        fmt.Println("request", req, time.Now())
    }

	// 允许短暂的并发请求，并同时保留总体速率限制

	// 带缓存 的速率限制器
    burstyLimiter := make(chan time.Time, 3)
	// 初始缓存 三个值，以便在任务开始时，可以并发的执行前三个任务
    for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now()
    }

    go func() {
		// 每 200ms 添加一个新的值到 burstyLimiter中 (满了会阻塞)
        for t := range time.Tick(200 * time.Millisecond) {
            burstyLimiter <- t
        }
    }()

    burstyRequests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)
    for req := range burstyRequests {
		// 初始三个请求会并发执行，后续 每隔 200 ms 执行一个任务
        <-burstyLimiter
        fmt.Println("request", req, time.Now())
    }
}