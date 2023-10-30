package main

import (
    "fmt"
	// strconv 包提供了数字解析能力
    "strconv"
)

func main() {

	// 解析为 浮点数，64 表示解析的数的位数
    f, _ := strconv.ParseFloat("1.234", 64)
    fmt.Println(f)

	// 解析为 整数， 0 表示自动推断字符串所表示的数字的进制，64 表示返回的整型数是以 64 位存储的
    i, _ := strconv.ParseInt("123", 0, 64)
    fmt.Println(i)

    d, _ := strconv.ParseInt("0x1c8", 0, 64)
    fmt.Println(d)

	// 解析为 无符号整型
    u, _ := strconv.ParseUint("789", 0, 64)
    fmt.Println(u)

	// 基础的 字符串 -> 十进制
    k, _ := strconv.Atoi("135")
    fmt.Println(k)

    _, e := strconv.Atoi("wat")
    fmt.Println(e)
}