package main

import (
    "errors"
    "fmt"
)

// 错误通常是最后一个返回值，并且是 error 类型，它是一个内建的接口
// type error interface {
// 	Error() string
// }
func f1(arg int) (int, error) {
    if arg == 42 {

		// 使用 errors.New() 构建一个 error 类型
        return -1, errors.New("can't work with 42")

    }

    return arg + 3, nil
}

// 自定义 错误类型
type argError struct {
    arg  int
    prob string
}
// 实现 error 接口 中的 Error() 方法
func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
    if arg == 42 {

        return -1, &argError{arg, "can't work with it"}
    }
    return arg + 3, nil
}

func main() {

    for _, i := range []int{7, 42} {
        if r, e := f1(i); e != nil {
            fmt.Println("f1 failed:", e)
        } else {
            fmt.Println("f1 worked:", r)
        }
    }
    for _, i := range []int{7, 42} {
        if r, e := f2(i); e != nil {
            fmt.Println("f2 failed:", e)
        } else {
            fmt.Println("f2 worked:", r)
        }
    }

    _, e := f2(42)
	// 通过 类型断言 来得到这个自定义错误类型的实例
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }
}