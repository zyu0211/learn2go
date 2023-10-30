package main

import (
    "fmt"
    "os"
)

func main() {

	// os.Args 提供原始命令行参数访问功能
    argsWithProg := os.Args
	// 切片中的第一个参数是该程序的路径，而 os.Args[1:]保存了程序全部的参数
    argsWithoutProg := os.Args[1:]

	// 使用标准的下标方式取得单个参数的值
    arg := os.Args[3]

    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)
}