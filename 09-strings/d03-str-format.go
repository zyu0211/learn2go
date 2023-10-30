package main

import (
    "fmt"
    "os"
)

type point struct {
    x, y int
}

func main() {

    p := point{1, 2}
	// %v: 常规值
    fmt.Printf("struct1: %v\n", p)

	// %+v: 如果值是结构体，%+v 的格式化输出内容包括其中的字段名
    fmt.Printf("struct2: %+v\n", p)

	// %#v: 会产生该值的源码片段
    fmt.Printf("struct3: %#v\n", p)

	// %T: 类型
    fmt.Printf("type: %T\n", p)

	// %t: 格式化 布尔值
    fmt.Printf("bool: %t\n", true)

	// %d: 十进制数字
    fmt.Printf("int: %d\n", 123)

	// %b: 二进制
    fmt.Printf("bin: %b\n", 14)

	// %c: 整数对于的字符
    fmt.Printf("char: %c\n", 33)

	// %x: 十六进制
    fmt.Printf("hex: %x\n", 456)

	// %f: 浮点数
    fmt.Printf("float1: %f\n", 78.9)

	// %e %E: 浮点数，以科学记数法表示
    fmt.Printf("float2: %e\n", 123400000.0)
    fmt.Printf("float3: %E\n", 123400000.0)

	// %s: 字符串 (转义字符会生效)
    fmt.Printf("str1: %s\n", "\"string\"")

	// %q: 字符串 (转义字符不会生效，类似 原始字符串)
    fmt.Printf("str2: %q\n", "\"string\"")

	// %x: 输出使用 base-16 编码的字符串， 每个字节使用 2 个字符表示
    fmt.Printf("str3: %x\n", "hex this")

	// %p: 指针
    fmt.Printf("pointer: %p\n", &p)

	// 带宽度的整数
    fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	// 带宽度的浮点数 (默认右对齐)
    fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	// 带宽度的浮点数，左对齐
    fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// 带宽度的字符串 (默认右对齐)
    fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	// 带宽度的字符串，左对齐 
    fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	// 格式化并返回一个字符串，但不打印输出
    s := fmt.Sprintf("sprintf: a %s", "string")
    fmt.Println(s)

	// 格式化并输出到 指定的 io 流 (Printf 是输出到 os.Stdout 的)
    fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}