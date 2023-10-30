package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

	// 在当前工作目录下，创建一个子目录
    err := os.Mkdir("subdir", 0755)
    check(err)

	// os.RemoveAll 会删除整个目录（类似于 rm -rf）
    defer os.RemoveAll("subdir")

	// 一个用于 创建临时文件 的帮助函数
    createEmptyFile := func(name string) {
        d := []byte("")
        check(os.WriteFile(name, d, 0644))
    }

    createEmptyFile("subdir/file1")

	// 创建一个有层级的目录 (类似于 mkdir -p)
    err = os.MkdirAll("subdir/parent/child", 0755)
    check(err)

    createEmptyFile("subdir/parent/file2")
    createEmptyFile("subdir/parent/file3")
    createEmptyFile("subdir/parent/child/file4")

	// ReadDir 列出目录的内容，返回一个 os.DirEntry 类型的切片对象
    c, err := os.ReadDir("subdir/parent")
    check(err)

    fmt.Println("Listing subdir/parent")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

	// Chdir 可以修改当前工作目录，类似于 cd
    err = os.Chdir("subdir/parent/child")
    check(err)

    c, err = os.ReadDir(".")
    check(err)

    fmt.Println("Listing subdir/parent/child")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

    err = os.Chdir("../../..")
    check(err)

    fmt.Println("Visiting subdir")
	// 遍历一个目录及其所有子目录
	// Walk 接受一个路径和回调函数，用于处理访问到的每个目录和文件
    err = filepath.Walk("subdir", visit)
}

// filepath.Walk 遍历访问到每一个目录和文件后，都会调用 visit
func visit(p string, info os.FileInfo, err error) error {
    if err != nil {
        return err
    }
    fmt.Println(" ", p, info.IsDir())
    return nil
}