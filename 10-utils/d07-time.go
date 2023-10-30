package main

import (
    "fmt"
    "time"
)

func main() {
    p := fmt.Println

	// 获取 当前时间
    now := time.Now()
    p(now)

	// 通过提供年月日等信息，构建一个 time
	// 时间总是与 Location (时区) 有关
    then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
    p(then)

	// 可以提取出时间的各个组成部分
    p(then.Year())
    p(then.Month())
    p(then.Day())
    p(then.Hour())
    p(then.Minute())
    p(then.Second())
    p(then.Nanosecond())
    p(then.Location())

    p(then.Weekday())

	// 比较两个时间的前后，精确到秒
    p(then.Before(now))
    p(then.After(now))
    p(then.Equal(now))

	// 方法 Sub 返回一个 Duration，表示 两个时间点的 间隔时间
    diff := now.Sub(then)
    p(diff)

	// 用各种单位来表示时间段的长度
    p(diff.Hours())
    p(diff.Minutes())
    p(diff.Seconds())
    p(diff.Nanoseconds())

	// 使用 Add 将时间后移一个时间段，或者使用一个 - 来将时间前移一个时间段
    p(then.Add(diff))
    p(then.Add(-diff))
}