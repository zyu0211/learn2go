package main

import (
    "fmt"
    "time"
)

// 工作者：用于处理任务
// 从 jobs 通道 中 获取任务
// 处理完成后，将结果写入到 通道 channel 中
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2
    }
}

func main() {

    const numJobs = 5
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

	// 启动 3 个 工作者
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

	// 向 任务通道 jobs 中添加任务
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
	
	// 关闭 任务 通道
    close(jobs)


	// 从 结果 通道 中取出结果
    for a := 1; a <= numJobs; a++ {
        <-results
    }
}