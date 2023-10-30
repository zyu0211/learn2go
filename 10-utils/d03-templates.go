package main

import (
    "os"
	// text/template 包为创建动态内容或向用户显示自定义输出提供了内置支持
    "text/template"
)

func main() {

	// 模板: 静态文本和包含在 ”动作“ 中用于动态插入内容的 混合体

	// 创建一个新模板
    t1 := template.New("t1")
	// 从字符串解析其正文
    t1, err := t1.Parse("Value is {{.}}\n")
    if err != nil {
        panic(err)
    }

	// 在 Parse 错误时返回 panic
    t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	// 执行模板，生成文本 
    t1.Execute(os.Stdout, "some text")
    t1.Execute(os.Stdout, 5)
    t1.Execute(os.Stdout, []string{
        "Go",
        "Rust",
        "C++",
        "C#",
    })

	// 辅助函数
    Create := func(name, t string) *template.Template {
        return template.Must(template.New(name).Parse(t))
    }

    t2 := Create("t2", "Name: {{.Name}}\n")

	// 对于 结构体 或 map，可以通过 {{.FieldName}} 动作来访问其字段
	// 匿名结构体
    t2.Execute(os.Stdout, struct {
        Name string
    }{"Jane Doe"})

    t2.Execute(os.Stdout, map[string]string{
        "Name": "Mickey Mouse",
    })

	// if/else 提供了条件执行模板
	// 如果一个值是类型的默认值，例如 0、空字符串、空指针等， 则该值被认为是 false
	// 下例中 使用 - 在动作中去除空格
    t3 := Create("t3",
        "{{if . -}} yes {{else -}} no {{end}}\n")
    t3.Execute(os.Stdout, "not empty")
    t3.Execute(os.Stdout, "")

	// range 块允许我们遍历切片，数组，映射或通道
	// 在 range 块内，{{.}} 动作被设置为迭代的当前项
    t4 := Create("t4",
        "Range: {{range .}}{{.}} {{end}}\n")
    t4.Execute(os.Stdout,
        []string{
            "Go",
            "Rust",
            "C++",
            "C#",
        })
}