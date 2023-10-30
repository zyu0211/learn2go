package main

import (
    "fmt"
    "sort"
)

func main() {

    strs := []string{"c", "a", "b"}
	// 对 string 类型 进行自然排序 (原地修改)
    sort.Strings(strs)
    fmt.Println("Strings:", strs)

    ints := []int{7, 2, 4}
	// 对 int 类型 进行自然排序
    sort.Ints(ints)
    fmt.Println("Ints:   ", ints)

	// 判断 一个 int 类型切片 是否有序
    s := sort.IntsAreSorted(ints)
    fmt.Println("Sorted: ", s)
}