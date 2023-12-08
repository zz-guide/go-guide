# go-guide
练习Go

## 工作空间
### basic
包含基础语法等练习，例如：字符串，切片，map等module

### advance
进阶，高级知识，例如：同步，锁，channel，并发等

### testing
单元测试等

### design
设计模式等

## 创建工作空间
```shell
cd ./basic && \
go work init
```

## 创建module
```shell
mkdir xxx && \
go mod init github.com/zz-guide/go-guide/basic/xxx
```

## module 加入工作空间
```shell
go work use ./xxx
```