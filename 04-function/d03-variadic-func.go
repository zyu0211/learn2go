package main

import "fmt"

// 变参函数：可以接收任意数量的 int 类型的参数
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {

    sum(1, 2)
    sum(1, 2, 3)

	// 传入 切片时：使用 func(slice...)
    nums := []int{1, 2, 3, 4}
    sum(nums...)
}