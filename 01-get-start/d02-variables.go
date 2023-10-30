package main

import "fmt"

func main() {

	// var 声名一个变量
	var a = "initial"
	fmt.Println(a)

	// 一次声名多个变量
	var b, c int = 1, 2
	fmt.Println(b, c)

	// 自动类型推断
	var d = true
	fmt.Println(d)

	// 声名但未初始化，会赋零值
	var e int
	fmt.Println(e)

	// ":=" 语法：声名并初始化（只能用于第一次出现的变量）
	f := "short"
	fmt.Println(f)
}
