package main

import (
    "os"

	// exec: 用其它（也许是非 Go）的进程，来完全替代当前的 Go 进程
    "os/exec"
    "syscall"
)

func main() {

	// 执行 ls 命令
	// Go 要求提供要执行的可执行文件的绝对路径，所以使用 exec.LookPath 找到它
	// 应该是 /bin/ls
    binary, lookErr := exec.LookPath("ls")
    if lookErr != nil {
        panic(lookErr)
    }

	// Exec 需要的参数是切片的形式的（不是放在一起的一个大字符串）
	// 注意，第一个参数需要是程序名
    args := []string{"ls", "-a", "-l", "-h"}

	// Exec 同样需要使用环境变量。这里我们仅提供当前的环境变量
    env := os.Environ()

	// syscall.Exec 调用
	// 如果这个调用成功，那么我们的进程将在这里结束，并被 /bin/ls -a -l -h 进程代替
	// 如果存在错误，那么我们将会得到一个返回值
    execErr := syscall.Exec(binary, args, env)
    if execErr != nil {
        panic(execErr)
    }
}