package main

import "fmt"

func main() {

    // 存放 5 个元素的 int 类型的数组
    var a [5]int
    fmt.Println("emp:", a)

    // 通过 索引下标 给数组中元素赋值和获取值
    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    // len() 返回素组长度
    fmt.Println("len:", len(a))

    // 声名并初始化
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

    // 二维数组
    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}