package main

import (
    "fmt"
	// 伪随机数 生成器
    "math/rand"
    "time"
)

func main() {

	// 返回一个随机的整数 n，且 0 <= n < 100
    fmt.Print(rand.Intn(100), ",")
    fmt.Print(rand.Intn(100))
    fmt.Println()

	// 返回一个64位浮点数 f，且 0.0 <= f < 1.0
    fmt.Println(rand.Float64())

    fmt.Print((rand.Float64()*5)+5, ",")
    fmt.Print((rand.Float64() * 5) + 5)
    fmt.Println()

	// 生成 一个 随机种子
	// 相同的随机种子，每次都会产生相同的随机数数字序列
	// 不同的随机种子，每次会产生不同的数字序列
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

    fmt.Print(r1.Intn(100), ",")
    fmt.Print(r1.Intn(100))
    fmt.Println()

    s2 := rand.NewSource(42)
    r2 := rand.New(s2)
    fmt.Print(r2.Intn(100), ",")
    fmt.Print(r2.Intn(100))
    fmt.Println()
    s3 := rand.NewSource(42)
    r3 := rand.New(s3)
    fmt.Print(r3.Intn(100), ",")
    fmt.Print(r3.Intn(100))
}