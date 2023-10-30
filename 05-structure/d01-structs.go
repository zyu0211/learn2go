package main

import "fmt"

// 定义一个结构体，包含两个字段
type person struct {
    name string
    age  int
}

// 函数：根据给定的名字，返回一个 person 结构体的指针
func newPerson(name string) *person {

    // 使用 结构体
    p := person{name: name}
    p.age = 42
    return &p
}

func main() {

    // 按序初始化结构体字段
    fmt.Println(person{"Bob", 20})

    // 初始化时，指定字段名称
    fmt.Println(person{name: "Alice", age: 30})

    // 未指定的字段，会被赋 默认值
    fmt.Println(person{name: "Fred"})

    fmt.Println(&person{name: "Ann", age: 40})

    fmt.Println(newPerson("Jon"))

    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    // 指针会自动解引用
    sp := &s
    fmt.Println(sp.age)

    sp.age = 51
    fmt.Println(sp.age)
}