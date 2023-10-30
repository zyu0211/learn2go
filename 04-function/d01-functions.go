package main

import "fmt"

// 定义一个函数：实现两数加法操作
func plus(a int, b int) int {

    return a + b
}

// 定义一个函数：实现三数加法操作
func plusPlus(a, b, c int) int {
    return a + b + c
}

func main() {

	// 调用函数
    res := plus(1, 2)
    fmt.Println("1+2 =", res)

    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)
}