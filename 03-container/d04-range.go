package main

import "fmt"

func main() {

	// range 遍历切片，进行求和
    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

	// 返回 索引 和 值
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

	// 遍历 map，迭代键值对
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

	// 遍历 map，只遍历键
    for k := range kvs {
        fmt.Println("key:", k)
    }

	// 在字符串中迭代 Unicode 码点 (i:起始字节位置，c:字符本身)
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}