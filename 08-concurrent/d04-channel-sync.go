package main

import (
    "fmt"
    "time"
)

// 使用 通道 来 同步 协程 之间的同步状态
func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    // worker 任务完成，向通道中发送一个同时
    done <- true
}

func main() {

    done := make(chan bool, 1)
    go worker(done)

    // 从通道中获取值
    // worker 任务完成之前，会一直阻塞
    <-done
}