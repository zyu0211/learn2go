package main

import (
    "fmt"
    "io"
    "os/exec"
)

func main() {

    // exec.Command 可以创建一个对象，来表示这个外部进程
    dateCmd := exec.Command("date")

    // .Output 是另一个帮助函数，常用于处理运行命令、等待命令完成并收集其输出
    // 如果没有错误，dateOut 将保存带有日期信息的字节
    dateOut, err := dateCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> date")
    fmt.Println(string(dateOut))

    // 示例：从外部进程的 stdin 输入数据并从 stdout 收集结果
    grepCmd := exec.Command("grep", "hello")

    // 明确的获取输入/输出管道，运行这个进程， 写入一些输入数据、读取输出结果，最后等待程序运行结束
    grepIn, _ := grepCmd.StdinPipe()
    grepOut, _ := grepCmd.StdoutPipe()
    grepCmd.Start()
    grepIn.Write([]byte("hello grep\ngoodbye grep"))
    grepIn.Close()
    grepBytes, _ := io.ReadAll(grepOut)
    grepCmd.Wait()

    fmt.Println("> grep hello")
    fmt.Println(string(grepBytes))

    // 注意，在生成命令时，我们需要提供一个明确描述命令和参数的数组，而不能只传递一个命令行字符串。
    // 如果你想使用一个字符串生成一个完整的命令，那么你可以使用 bash 命令的 -c 选项
    lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
    lsOut, err := lsCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> ls -a -l -h")
    fmt.Println(string(lsOut))
}