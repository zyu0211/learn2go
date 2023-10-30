package main

import (
    "fmt"
    "os"
)

func main() {

	// defer 用于确保程序在执行完成后，会调用某个函数
	// 类似于 finally

    f := createFile("/tmp/defer.txt")
	// 在程序结束时，执行此函数，关闭该文件
    defer closeFile(f)
    writeFile(f)
}

func createFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}

func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {
    fmt.Println("closing")
    err := f.Close()

    // 关闭文件时，进行错误检查
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}