package main

import (
    "fmt"
    "unicode/utf8"
)

// Go 中，字符串是一个只读的 byte 类型的切片
// Go 中，字符的概念被称为 rune (表示 Unicode 编码的整数)

func main() {

	// 字符串 s，是一个泰语中的单词
    const s = "สวัสดี"

	// len() 返回的是 byte 切片的长度
    fmt.Println("Len:", len(s))

	// 在每个索引处生成原始字节值
    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
    fmt.Println()

	// 计算字符串中有多少 rune
    fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// 遍历获取 每个 rune 及其在字符串中的偏移量
    for idx, runeValue := range s {
        fmt.Printf("%#U starts at %d\n", runeValue, idx)
    }

	// 手动实现对 string 的遍历
    fmt.Println("\nUsing DecodeRuneInString")
    for i, w := 0, 0; i < len(s); i += w {
        runeValue, width := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%#U starts at %d\n", runeValue, i)
        w = width

        examineRune(runeValue)
    }
}

func examineRune(r rune) {

    if r == 't' {
        fmt.Println("found tee")
    } else if r == 'ส' {
        fmt.Println("found so sua")
    }
}