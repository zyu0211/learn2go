package main

import "fmt"

func main() {

	// 基本用法
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// 只要 if，不要 else
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// 在 if 的条件判断前，声名语句, 用 分号 隔开（可以在 if 的所有条件分支中使用）
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
