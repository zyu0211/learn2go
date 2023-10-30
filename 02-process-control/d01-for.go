package main

import "fmt"

func main() {

	// 最基础的方式，单个循环条件
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// 经典用法：for-初始值-条件判断-后续处理
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// 死循环
	for {
		fmt.Println("loop")
		break
	}

	// 使用 continue
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
