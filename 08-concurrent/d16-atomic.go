package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {

	// 无符号整数，表示计数器
    var ops uint64

	// wait group，用于等待所有协程的完成
    var wg sync.WaitGroup

    for i := 0; i < 50; i++ {
        wg.Add(1)

        go func() {
            for c := 0; c < 1000; c++ {

				// AddUint64 来让计数器自动增加
				// 需要使用 & 语法给定 ops 的内存地址
                atomic.AddUint64(&ops, 1)
            }
            wg.Done()
        }()
    }

    wg.Wait()

	// 可以使用 atomic.LoadUint64 之类的函数，在原子更新的同时安全地读取它们
    fmt.Println("ops:", ops)
}