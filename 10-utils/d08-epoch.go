package main

import (
    "fmt"
    "time"
)

func main() {

    now := time.Now()
	// 秒数
	secs := now.Unix()
	// 纳秒数
    nanos := now.UnixNano()
    fmt.Println(now)

	// 毫秒数，需手动计算
    millis := nanos / 1000000
    fmt.Println(secs)
    fmt.Println(millis)
    fmt.Println(nanos)

	// 将 整数秒或纳秒 转化到相应的时间
    fmt.Println(time.Unix(secs, 0))
    fmt.Println(time.Unix(0, nanos))
}