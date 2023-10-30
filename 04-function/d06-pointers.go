package main

import "fmt"

// 值传递
func zeroval(ival int) {
    ival = 0
}

// 引用传递
func zeroptr(iptr *int) {
	// 对解引用的指针赋值
    *iptr = 0
}

func main() {
    i := 1
    fmt.Println("initial:", i)

	// 值传递 不能 改变外部 i
    zeroval(i)
    fmt.Println("zeroval:", i)

	// 引用传递 能 改变外部 i
    zeroptr(&i)
    fmt.Println("zeroptr:", i)

    fmt.Println("pointer:", &i)
}