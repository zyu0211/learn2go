package main

import "os"

func main() {

	// panic：用來 抛出 程序正常运行中不应该出现的错误，或者我们不准备优雅处理的错误。
	// panic 通常会输出一个错误消息和协程追踪信息，并以非零的状态退出程序
	// 类似于 throw

	// 手动抛出错误 
    panic("a problem")

    _, err := os.Create("/tmp/file")
    if err != nil {
		// 创建新文件时遇到意外错误，使用 panic 抛出此錯誤
        panic(err)
    }
}