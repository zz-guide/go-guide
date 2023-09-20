package protocol

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
)

// 采用ding

// Encode 将消息编码
func Encode(message string) ([]byte, error) {
	// 读取消息的长度，转换成int32类型（占4个字节）
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	// 写入消息头
	//err := binary.Write(pkg, binary.LittleEndian, length)
	err := binary.Write(pkg, binary.BigEndian, length)
	if err != nil {
		return nil, err
	}
	// 写入消息实体
	//err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	err = binary.Write(pkg, binary.BigEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

// Decode 解码消息
func Decode(reader *bufio.Reader) (string, error) {
	// 读取消息的长度
	lengthByte, _ := reader.Peek(4) // 读取前4个字节的数据
	fmt.Println("lengthByte:", lengthByte, string(lengthByte))
	lengthBuff := bytes.NewBuffer(lengthByte)
	// 实际内容的长度
	var length int32
	//err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	// 按照大端序读取数据包的长度
	err := binary.Read(lengthBuff, binary.BigEndian, &length)
	if err != nil {
		return "", err
	}

	fmt.Println("ssss-length:", length)
	fmt.Println("lengthBuff:", lengthBuff.String())
	fmt.Println("reader.Buffered():", reader.Buffered())
	// Buffered返回缓冲中现有的可读取的字节数。
	// Buffered中的数据如果不满足一个数据包大小，说明读取完毕了，直接返回
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	// 读取一个数据包大小的数据
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
