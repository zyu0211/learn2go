package main

import (
    "fmt"
	// 标准库的 strings 包提供了很多有用的字符串相关的函数
    s "strings"
)

// 给 fmt.Println 一个别名（以便后续的使用）
var p = fmt.Println

func main() {

	// 以下是一些 strings 包中的例子
	// 由于它们都是包的函数，而不是字符串对象自身的方法，
	// 这意味着在调用函数时，需要将字符串作为第一个参数进行传递

	// 1、是否包含子串
    p("Contains:  ", s.Contains("test", "es"))

	// 2、统计 字符 出现的次数
    p("Count:     ", s.Count("test", "t"))

	// 3、是否为指定前缀
    p("HasPrefix: ", s.HasPrefix("test", "te"))

	// 4、是否为指定后缀
    p("HasSuffix: ", s.HasSuffix("test", "st"))
	
	// 5、字符所在的索引位置
    p("Index:     ", s.Index("test", "e"))

	// 6、字符串数组 以 指定的符号进行连接
    p("Join:      ", s.Join([]string{"a", "b"}, "-"))

	// 7、将字符串 重复 n 次
    p("Repeat:    ", s.Repeat("a", 5))

	// 8、将所有的 “o”，替换为 “0” 
    p("Replace:   ", s.Replace("foo", "o", "0", -1))

	// 9、将 第一个 “o”，替换为 “0”
    p("Replace:   ", s.Replace("foo", "o", "0", 1))

	// 10、字符串分割
    p("Split:     ", s.Split("a-b-c-d-e", "-"))

	// 11、转为小写
    p("ToLower:   ", s.ToLower("TEST"))

	// 12、转为大写
    p("ToUpper:   ", s.ToUpper("test"))
    p()

    p("Len: ", len("hello"))
    p("Char:", "hello"[1])
}