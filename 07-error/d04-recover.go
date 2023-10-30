package main

import "fmt"


// recover 可以从 panic 中 恢复，阻止 panic 中止程序，并让它继续执行。
// 类似于 catch

func mayPanic() {
    panic("a problem")
}

func main() {

	// 必须在 defer 函数中调用 recover
    defer func() {
		// recover() 的返回值是在调用 panic 时抛出的错误
        if r := recover(); r != nil {

            fmt.Println("Recovered. Error:\n", r)
        }
    }()

	// 当跳出引发 panic 的函数时，defer 会被激活，其中的 recover 会捕获 panic
    mayPanic()

	// 此处的代码不会执行
    fmt.Println("After mayPanic()")
}