package main

import "fmt"

func main() {

	// 创建长度为 3，string类型的 空切片
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// 切片赋值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// 切片长度
	fmt.Println("len:", len(s))

	// 添加元素
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// 切片拷贝
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// "切片" 操作
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	// 声名并初始化
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// slices 组成多维数据结构
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
