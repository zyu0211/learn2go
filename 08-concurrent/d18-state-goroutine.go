package main

import (
    "fmt"
    "math/rand"
    "sync/atomic"
    "time"
)

type readOp struct {
    key  int
    resp chan int
}
type writeOp struct {
    key  int
    val  int
    resp chan bool
}

func main() {

	// 共享变量
	// 用来计算 读/写 操作的次数
    var readOps uint64
    var writeOps uint64

	// 不带缓存的两个通道
	// 基于通道的阻塞，实现对共享变量的 互斥操作
    reads := make(chan readOp)
    writes := make(chan writeOp)

    go func() {
        var state = make(map[int]int)
        for {
            select {
            case read := <-reads:
                read.resp <- state[read.key]
            case write := <-writes:
                state[write.key] = write.val
                write.resp <- true
            }
        }
    }()

    for r := 0; r < 100; r++ {
        go func() {
            for {
                read := readOp{
                    key:  rand.Intn(5),
                    resp: make(chan int)} // 非缓存通道
			    // 向 reads 通道中 发送 一个数据
                reads <- read
				// 阻塞，等待读操作完成
                <-read.resp
				// 原子修改 读操作 的 计数
                atomic.AddUint64(&readOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    for w := 0; w < 10; w++ {
        go func() {
            for {
                write := writeOp{
                    key:  rand.Intn(5),
                    val:  rand.Intn(100),
                    resp: make(chan bool)} // 非缓存通道
				// 向 writes 通道中 发送 一个数据
                writes <- write
				// 阻塞，等待 写操作 完成
                <-write.resp
				// 原子修改 写操作 的 计数
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

	// main() 会等待 1 s
	// 即：会统计 1 s 时间内有多少次 读/写 操作
    time.Sleep(time.Second)

	// 原子操作，安全的去读 写操作计数 的值
    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:", readOpsFinal)

	// 原子操作，安全的去读 读操作计数 的值
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)
}