package main

import (
    "fmt"
    "sync"
)

// 需求：在多个 goroutine 同时更新 Container 中的 counters
// 保证数据的安全性，需要添加了一个 互斥锁 Mutex 来同步访问
// 注意：不能复制互斥锁，如果需要传递这个 struct，需要使用指针
type Container struct {
    mu       sync.Mutex
    counters map[string]int
}

func (c *Container) inc(name string) {

	// 上锁
    c.mu.Lock()
	// 释放锁
    defer c.mu.Unlock()
    c.counters[name]++
}

func main() {
	
	// 互斥量的零值是可用的
	// 声名 Container 时，不需要初始化 Mutex
    c := Container{
        counters: map[string]int{"a": 0, "b": 0},
    }

    var wg sync.WaitGroup

	// 函数：在循环中递增对 name 的计数
    doIncrement := func(name string, n int) {
        for i := 0; i < n; i++ {
            c.inc(name)
        }
        wg.Done()
    }

    wg.Add(3)
    go doIncrement("a", 10000)
    go doIncrement("a", 10000)
    go doIncrement("b", 10000)

    wg.Wait()
    fmt.Println(c.counters)
}