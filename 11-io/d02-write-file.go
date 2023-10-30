package main

import (
    "bufio"
    "fmt"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

	// 直接 写入
    d1 := []byte("hello\ngo\n")
    err := os.WriteFile("/tmp/text1.txt", d1, 0644)
    check(err)

	// 更细粒度的写入
	
	// 新建一个文件
    f, err := os.Create("/tmp/text2.txt")
    check(err)

    defer f.Close()

	// 基于 字节切片 写入
    d2 := []byte{115, 111, 109, 101, 10}
    n2, err := f.Write(d2)
    check(err)
    fmt.Printf("write %d bytes\n", n2)

	// 写入字符串
    n3, err := f.WriteString("writes\n")
    check(err)
    fmt.Printf("write %d bytes\n", n3)

	// 将缓冲区的数据写入硬盘
    f.Sync()

	// bufio 还提供了的带缓冲的 Writer
    w := bufio.NewWriter(f)
    n4, err := w.WriteString("buffered\n")
    check(err)
    fmt.Printf("write %d bytes\n", n4)

	// 使用 Flush 来确保，已将所有的缓冲操作应用于底层 writer
    w.Flush()

}