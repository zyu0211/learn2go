package main

import (
    "fmt"
	// filepath 包为 文件路径 ，提供了方便的 跨操作系统 的解析和构建函数
    "path/filepath"
    "strings"
)

func main() {

	// join 参照传入顺序构造一个对应层次结构的路径
    p := filepath.Join("dir1", "dir2", "filename")
    fmt.Println("p:", p)

	// Join 会删除多余的分隔符和目录，使得路径更加规范
    fmt.Println(filepath.Join("dir1//", "filename"))
    fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// Dir 和 Base 用于分割路径中的目录和文件
	// Split 可以一次性实现 上面两个函数的结果
    fmt.Println("Dir(p):", filepath.Dir(p))
    fmt.Println("Base(p):", filepath.Base(p))

	// 是否为 绝对路径
    fmt.Println(filepath.IsAbs("dir/file"))
    fmt.Println(filepath.IsAbs("/dir/file"))

    filename := "config.json"

	// 分割 出 文件 拓展名
    ext := filepath.Ext(filename)
    fmt.Println(ext)

	// 想获取 文件名清除扩展名后 的值
    fmt.Println(strings.TrimSuffix(filename, ext))

	// Rel 寻找 basepath 与 targpath 之间的相对路径
	// 如果相对路径不存在，则返回错误
    rel, err := filepath.Rel("a/b", "a/b/t/file")
    if err != nil {
        panic(err)
    }
    fmt.Println(rel)

    rel, err = filepath.Rel("a/b", "a/c/t/file")
    if err != nil {
        panic(err)
    }
    fmt.Println(rel)
}