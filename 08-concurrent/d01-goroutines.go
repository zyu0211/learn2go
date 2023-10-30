package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// 同步直接调用
	f("direct")

	// 使用 Go 协程，并发执行
	go f("goroutine")

	// 为 匿名函数 启动一个协程
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 等待 1 s，待 异步 执行的两个协程 执行完成
	// 可以使用 WaitGroup 进行等待，后面会讲解
	time.Sleep(time.Second)
	fmt.Println("done")
}
