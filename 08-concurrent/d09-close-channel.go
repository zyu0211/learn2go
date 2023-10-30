package main

import "fmt"

func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    go func() {
        for {
			// 通道已经关闭，并且通道中所有的值都已经接收完毕，那么 more 的值将是 false
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }

	// 关闭一个通道 意味着不能再向这个通道发送值
	// 但 还可以 读取 通道中 已经存在的值
    close(jobs)
    fmt.Println("sent all jobs")

    <-done
}