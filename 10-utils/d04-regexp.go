package main

import (
    "bytes"
    "fmt"
	// regexp 提供对正则表达式的支持
    "regexp"
)

func main() {
	// 最简单 的 用法
	// 测试一个字符串是否符合一个表达式
    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    fmt.Println(match)

	// 进阶用法
	// 通过 Compile 得到一个优化过的 Regexp 结构体
    r, _ := regexp.Compile("p([a-z]+)ch")

	// 判断一个字符串是否匹配
    fmt.Println(r.MatchString("peach"))

	// 查找匹配的字符串
    fmt.Println(r.FindString("peach punch"))

	// 查找首次匹配的字符串，返回匹配 开始 和 结束 位置的索引
    fmt.Println("idx:", r.FindStringIndex("peach punch"))

	// 返回完全匹配和局部匹配的字符串
    fmt.Println(r.FindStringSubmatch("peach punch"))

	// 返回完全匹配和局部匹配位置的索引
    fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// 返回 全部的 匹配项
    fmt.Println(r.FindAllString("peach punch pinch", -1))

	// 返回全 部的 完全匹配和局部匹配位置的索引
    fmt.Println("all:", r.FindAllStringSubmatchIndex(
        "peach punch pinch", -1))

	// 返回 前 2 个 匹配项
    fmt.Println(r.FindAllString("peach punch pinch", 2))

	// 输入为 []byte 类型
    fmt.Println(r.Match([]byte("peach")))

	// 创建正则表达式的全局变量时，也可以使用 MustCompile
	// MustCompile 用 panic 代替返回一个错误 ，这样使用全局变量更加安全
    r = regexp.MustCompile("p([a-z]+)ch")
    fmt.Println("regexp:", r)

	// 替换 匹配 的子串
    fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

    in := []byte("a peach")
	// Func 变体允许您使用给定的函数来转换匹配的文本
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))
}