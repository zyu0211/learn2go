package main

import "fmt"

// 闭包: 定义一个函数，其返回值是一个函数
// 返回的函数使用闭包的方式，隐藏变量 i。
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {

	// 每次调用都会更新 i 的值
    nextInt := intSeq()

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    newInts := intSeq()
    fmt.Println(newInts())
}