package main

import (
    "fmt"
    "time"
)

func main() {
    p := fmt.Println

    t := time.Now()
	// 遵循 RFC3339， 并使用对应的 布局（layout）常量进行格式化
    p(t.Format(time.RFC3339))

	// 时间解析使用与 Format 相同的布局值
    t1, e := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
    p(t1)

	// 自定义布局
    p(t.Format("3:04PM"))
    p(t.Format("Mon Jan _2 15:04:05 2006"))
    p(t.Format("2006-01-02T15:04:05.999999-07:00"))
    form := "3 04 PM"
    t2, e := time.Parse(form, "8 41 PM")
    p(t2)

    fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
        t.Year(), t.Month(), t.Day(),
        t.Hour(), t.Minute(), t.Second())

	// 当输入的时间格式不正确时，Parse 会返回一个解析错误
    ansic := "Mon Jan _2 15:04:05 2006"
    _, e = time.Parse(ansic, "8:41PM")
    p(e)
}