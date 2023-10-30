package main

import "fmt"

func main() {

	// 创建一个 空 map
    m := make(map[string]int)

	// 设置键值对
    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map:", m)

	// 获取键的值
    v1 := m["k1"]
    fmt.Println("v1: ", v1)

	// 返回键值对数量
    fmt.Println("len:", len(m))

	// 从 map 中删除一个键值对
    delete(m, "k2")
    fmt.Println("map:", m)

	// 使用第二个返回值，进行判空
    _, prs := m["k2"]
    fmt.Println("prs:", prs)

	// 声名并初始化
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}