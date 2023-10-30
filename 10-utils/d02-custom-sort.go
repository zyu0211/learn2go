package main

import (
    "fmt"
    "sort"
)

// 需求：实现一个定制排序，即：按字符串的长度排序

// 1、创建 []string 的类型别名 (自定义类型)
type byLength []string

// 2、自定义类型 实现 sort.Interface 中的三个方法： Len()、Swap()、Less()
func (s byLength) Len() int {
    return len(s)
}
func (s byLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {	// Less 将控制实际的自定义排序逻辑
    return len(s[i]) < len(s[j])
}

func main() {
    fruits := []string{"peach", "banana", "kiwi"}

	// sort.Sort 来实现自定义排序
	// 需要将切片 fruits 强转为 byLength 类型
    sort.Sort(byLength(fruits))

    fmt.Println(fruits)
}