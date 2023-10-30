package main

import (
    "fmt"
    "os"
)

func main() {

    defer fmt.Println("!")

	// os.Exit 可以立即以给定的状态退出程序
	// 当使用 os.Exit 时 defer 将不会 被执行，所以这里的 fmt.Println 将永远不会被调用
    os.Exit(3)
}