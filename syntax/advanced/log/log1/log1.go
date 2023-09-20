package logger

import (
	"log"
	"os"
)

var baseLogger *log.Logger

func TestBaseLog() {
	file, err := os.OpenFile("test.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		log.Fatalln("fail to create test.log file!", err)
	}

	defer file.Close()
	baseLogger = log.New(file, "Debug:", log.LstdFlags|log.Lshortfile) // 日志文件格式:log包含时间及文件行数
	//baseLogger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	baseLogger.Println("你好")
	baseLogger.Println("张三")
}
