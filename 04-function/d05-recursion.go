package main

import "fmt"

// n 不为 0 时，函数一直递归调用自身
func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

func main() {
    fmt.Println(fact(7))

	// 闭包的递归
	// 要求：在定义闭包前，用类型化的 var 显示的声名闭包
    var fib func(n int) int

	// 闭包
    fib = func(n int) int {
        if n < 2 {
            return n
        }
        return fib(n-1) + fib(n-2)

    }

    fmt.Println(fib(7))
}