package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// 行过滤器（line filter） 是一种常见的程序类型， 
// 它读取 stdin 上的输入，对其进行处理，然后将处理结果打印到 stdout。
// grep 和 sed 就是常见的行过滤器

func main() {

	// 用带缓冲的 scanner 包装无缓冲的 os.Stdin， 
	// 这为我们提供了一种方便的 Scan 方法， 将 scanner 前进到下一个 令牌（默认为：下一行）
    scanner := bufio.NewScanner(os.Stdin)
    
	for scanner.Scan() {

		// Text 返回当前的 token
        ucl := strings.ToUpper(scanner.Text())

        fmt.Println(ucl)
    }

	// 检查 Scan 的错误
	// 文件结束符（EOF）是可以接受的，它不会被 Scan 当作一个错误
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
}