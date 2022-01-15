package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	//1.exec.LookPath方法用于在操作系统查找符合的可以执行的二进制文件
	//如果只传入名字，会进行查找，如果传入的就是路径，则不会进行查找
	//结果可能是一个绝对路径，或者相对于当前目录的路径
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	fmt.Println(binary)

	//设置要执行的命令和参数
	args := []string{"ls", "-a", "-l", "-h"}
	//设置要在那个环境执行，其实就是设置当前环境变量
	env := os.Environ()

	//正真执行系统命令
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

	//执行结果会自动在控制台输出
}
