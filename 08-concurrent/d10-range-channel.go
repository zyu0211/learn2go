package main

import "fmt"

func main() {

    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)

	// range 迭代从 通道 queue 中得到每个值 
	// 因为 close 通道 queue，所以，迭代 2 个值之后会结束
    for elem := range queue {
        fmt.Println(elem)
    }
}