package main

import "fmt"

// 结构体
type rect struct {
    width, height int
}

// area 方法 拥有 *rect 类型接收器(receiver)
func (r *rect) area() int {
    return r.width * r.height
}

func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

func main() {
	// 创建结构体
    r := rect{width: 10, height: 5}

	// 使用 结构体 调用 方法
    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())

	// 调用方法时，Go 会自动处理 值 与 指针 之间的转换
    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
}