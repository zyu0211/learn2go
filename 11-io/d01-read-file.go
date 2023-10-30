package main

import (
	// bufio 包实现了一个缓冲读取器，这可能有助于提高许多小读操作的效率，以及它提供了很多附加的读取函数
    "bufio"
    "fmt"
    "io"
    "os"
)

// 读取文件需要经常进行错误检查
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

	// 最基本的文件读取任务
	// 直接将文件内容读取到内存中
    dat, err := os.ReadFile("/tmp/test.txt")
    check(err)
    fmt.Print(string(dat))

	// 更可控的进行文件读取

	// 打开一个文件
    f, err := os.Open("/tmp/test.txt")
    check(err)

	// 每次 读取 5 个字节
    b1 := make([]byte, 5)
	// n1 为实际读取的字节数
    n1, err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// 直接跳到文件中的某个位置 进行读取
    o2, err := f.Seek(6, 0)
    check(err)
    b2 := make([]byte, 2)
    n2, err := f.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d: ", n2, o2)
    fmt.Printf("%v\n", string(b2[:n2]))

    o3, err := f.Seek(6, 0)
    check(err)
    b3 := make([]byte, 2)
	// io 包提供了一个更健壮的实现 ReadAtLeast，用于读取上面那种文件
    n3, err := io.ReadAtLeast(f, b3, 2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// 回到文件开头
    _, err = f.Seek(0, 0)
    check(err)

    r4 := bufio.NewReader(f)
    b4, err := r4.Peek(5)
    check(err)
    fmt.Printf("5 bytes: %s\n", string(b4))

	// 任务结束后要关闭这个文件 （通常这个操作应该在 Open 操作后立即使用 defer 来完成）
    f.Close()
}