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

	// 创建临时文件最简单的方法是调用 os.CreateTemp 函数。它会创建并打开文件，我们可以对文件进行读写。 
	// 第一个参数，创建该文件的路径，""表示系统默认，在类 Unix 操作系统下，临时目录一般是 /tmp
	// 第二个参数，生成的文件的名称的前缀
    f, err := os.CreateTemp("", "sample")
    check(err)

	// 打印临时文件的名称
	// 文件名以 os.CreateTemp 函数的第二个参数作为前缀， 剩余的部分会自动生成，以确保并发调用时，生成不重复的文件名。
    fmt.Println("Temp file name:", f.Name())

	// defer 删除该文件
	// 尽管操作系统会自动在某个时间清理临时文件，但手动清理是一个好习惯
    defer os.Remove(f.Name())

    _, err = f.Write([]byte{1, 2, 3, 4})
    check(err)

	// MkdirTemp 创建一个临时 目录
	// 参数与 CreateTemp 相同，但是它返回的是一个 目录名，而不是一个打开的文件
    dname, err := os.MkdirTemp("", "sampledir")
    fmt.Println("Temp dir name:", dname)

    defer os.RemoveAll(dname)

    fname := filepath.Join(dname, "file1")
    err = os.WriteFile(fname, []byte{1, 2}, 0666)
    check(err)
}