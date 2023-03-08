######源码调试器
1.delve
2.gdb
3.lldb

4.go tool compile

######dlv调试
break,b
continue,c
next
step


#######go build 

生成汇编代码：.o文件

方法一：
go tool compile -N -l -S main.go

方法二：
go build -gcflags "-N -l" -S main.go

方法三: go tool objdump
首先先编译程序: go tool compile -N -l once.go,
使用go tool objdump once.o反汇编出代码
(或者使用go tool objdump -s Do once.o反汇编特定的函数：)
