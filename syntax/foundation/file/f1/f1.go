package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	ReadBuf()
}

func Pwd() {
	path, _ := os.Getwd()
	fmt.Println("当前路径:", path)
}

func mkdir() {
	path, _ := os.Getwd()
	err := os.Mkdir(path+"/_log", 0777)
	if err != nil {
		//单层目录创建，重复的时候会报错
		log.Println(err)
	}

	/*err = os.MkdirAll("D:\\theWayGo\\src\\gotest\\base-package\\os\\dir2\\test", 0777)
	if err != nil {
		//多级目录创建，会安目录一级级创建，但同样的路径，删除的时候要注意路径，删除只会针对路径最后一级，重复创建的时候不会报错
		log.Println(err)
	}*/
}

func createFile() {
	path, _ := os.Getwd()
	file, err := os.Create(path + "/_log/read.txt")
	if err != nil {
		log.Println(err.Error())
	}

	defer file.Close()
	fmt.Println(file.Name())
}

//使用ioutil包进行文件写入
func ioutilWrite() {
	path, _ := os.Getwd()
	content := "大道至简1\n大道至简2\n"
	err := ioutil.WriteFile(path+"/_log/read.txt", []byte(content), 0777)
	if err != nil {
		log.Println("write fail")
		panic(err.Error())
	}
}

//使用os.OpenFile相关函数进行文件写入
func openWrite() {
	path, _ := os.Getwd()
	file, err := os.OpenFile(path+"/_log/read.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Println(err.Error())
		panic(err.Error())
	}

	defer file.Close()

	_, _ = file.Seek(0, 2)
	content := []byte("Go 会是转折点么？")
	_, err = file.Write(content)
	if err != nil {
		log.Println(err.Error())
		panic(err.Error())
	}
}

//使用bufio包中的相关函数写入文件
func bufioWrite() {
	path, _ := os.Getwd()
	file, err := os.OpenFile(path+"/_log/read.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Println(err.Error())
		panic(err.Error())
	}

	defer file.Close()

	file.Seek(0, io.SeekEnd) // 代表从末尾开始写
	content := []byte("但行好事莫问前程，自顾向前何愁天不怜")

	//使用NewWriter方法返回的io.Writer缓冲默认大小为4096，也可以使用NewWriterSize方法设置缓存的大小
	buf := bufio.NewWriter(file)

	//将文件写入缓存
	_, err = buf.Write(content)
	if err != nil {
		log.Println(err.Error())
		panic(err.Error())
	}

	//从缓存写入到文件中
	if err = buf.Flush(); err != nil {
		log.Println(err.Error())
		panic(err.Error())
	}
}

// ReadFile 直接读
func ReadFile() {
	path, _ := os.Getwd()
	name := path + "/_log/read.txt"
	if contents, err := ioutil.ReadFile(name); err == nil {
		fmt.Println("ReadFile:", string(contents))
	}
}

// 先从文件读取到file中，在从file读取到buf, buf在追加到最终的[]byte
func ReadBuf() {
	path, _ := os.Getwd()
	name := path + "/_log/read.txt"

	//获得一个file
	f, err := os.Open(name)
	if err != nil {
		fmt.Println("read fail")
		return
	}

	//把file读取到缓冲区中
	defer f.Close()
	var chunk []byte
	buf := make([]byte, 1024)

	for {
		//从file读取到buf中
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("read buf fail", err)
			return
		}
		//说明读取结束
		if n == 0 {
			break
		}
		//读取到最终的缓冲区中
		chunk = append(chunk, buf[:n]...)
	}

	fmt.Println(string(chunk))
}

//先从文件读取到file, 在从file读取到Reader中，从Reader读取到buf, buf最终追加到[]byte，这个排第三
func ReadReader() {
	path, _ := os.Getwd()
	name := path + "/_log/read.txt"

	fi, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	r := bufio.NewReader(fi)
	var chunks []byte

	buf := make([]byte, 1024)

	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		//fmt.Println(stringSearch(buf))
		chunks = append(chunks, buf...)
	}
	fmt.Println(string(chunks))
}

//ReadReadAll 读取到file中，再利用ioutil将file直接读取到[]byte中, 这是最优
func ReadReadAll() {
	path, _ := os.Getwd()
	name := path + "/_log/read.txt"
	f, err := os.Open(name)
	if err != nil {
		fmt.Println("read file fail", err)
		return
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read to fd fail", err)
		return
	}

	fmt.Println(string(fd))
}
