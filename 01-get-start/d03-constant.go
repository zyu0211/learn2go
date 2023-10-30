package main

import (
	"fmt"
	"math"
)

// const 用于声名一个常量
const s string = "constant"

func main() {
	fmt.Println(s)

	// const 语句可以出现在任何 var 语句可以出现的地方
	const n = 500000000

	// 常量表达式 可以执行任意精度的运算
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
