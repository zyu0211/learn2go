package main

import "fmt"

// 基础 结构体
type base struct {
    num int
}
// base 的一个方法
func (b base) describe() string {
    return fmt.Sprintf("base with num=%v", b.num)
}

// 拓展 结构体：内部嵌入了一个 基础 结构体
type container struct {
    base
    str string
}

func main() {

    co := container{
        base: base{
            num: 1,
        },
        str: "some name",
    }

	// co 直接访问 base 中定义的字段
    fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

    fmt.Println("also num:", co.base.num)

	// co 直接调用 base 的方法
    fmt.Println("describe:", co.describe())

    type describer interface {
        describe() string
    }

	// container 嵌入了 base ，所以 container 也实现了 describer 接口
    var d describer = co
    fmt.Println("describer:", d.describe())
}