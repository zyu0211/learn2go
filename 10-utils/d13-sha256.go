package main

import (
	// Go 在多个 crypto/* 包中实现了一系列散列函数
	// sha256 散列 (hash)
    "crypto/sha256"
    "fmt"
)

func main() {
    s := "sha256 this string"

	// 新的散列函数
    h := sha256.New()

	// 写入要处理的字节
    h.Write([]byte(s))

	// Sum 得到最终的散列值的字符切片
	// Sum 接收一个参数，可以用来给现有的字符切片追加额外的字节切片，但是一般都不需要这样做
    bs := h.Sum(nil)

	// SHA256 值经常以 16 进制输出
    fmt.Println(s)
    fmt.Printf("%x\n", bs)
}