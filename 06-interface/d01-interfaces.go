package main

import (
    "fmt"
    "math"
)

// 定义 一个 几何体 的基本接口
type geometry interface {
    area() float64
    perim() float64
}

// 结构体：矩形
type rect struct {
    width, height float64
}
// 结构体：圆形
type circle struct {
    radius float64
}

// 在 Go 中实现一个接口，我们只需要实现接口中的所有方法

// rect 的方法，实现 接口 geometry 中 的方法
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

// circle 的方法，实现 接口 geometry 中 的方法
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

// 根据传入的接口类型 的不同实现，以调用不同的方法
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    measure(r)
    measure(c)
}