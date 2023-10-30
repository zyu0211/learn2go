package main

import "fmt"

// 定义一和个函数：返回 两个 int 类型的值
func vals() (int, int) {
    return 3, 7
}

func main() {

	// 接收两个返回
    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

	// 只接收 其中一个返回值
    _, c := vals()
    fmt.Println(c)
}