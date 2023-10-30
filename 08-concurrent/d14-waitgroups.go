package main

import (
    "fmt"
    "sync"
    "time"
)

// 工作者，用于在 协程 中 处理任务
func worker(id int) {
    fmt.Printf("Worker %d starting\n", id)

    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {

	// 声名一个 wait group
	// 如果 wg 要显示传递到函数中，需要使用 指针
    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
		// Add(1): 需要等待的任务数量 加一
        wg.Add(1)

        i := i

        go func() {
			// Done(): 完成一个任务
            defer wg.Done()
            worker(i)
        }()
    }

	// main 中 会一直等待，直到 wait group 中 等待的任务 全部完成
    wg.Wait()

}