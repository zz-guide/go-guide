/**
作用：用于将excel内容转化成json并输出到文件，其他语言可通过读取该文件内容进行便利操作。
场景：第一行是表头，表名都有哪些字段，从第2行开始为正文内容
	 目前只支持一个sheet
	 超过表头列的内容不会读取
*/
package read

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/xuri/excelize/v2"
)

const SheetName = "考纲词库介绍"
const FilePath = "/Users/xulei/Downloads/valid.xlsx"
const OutFilePath = "/Users/xulei/Downloads/读取json.txt"

var mKey = []string{"stage", "material", "book", "valid"}
var mKeyLen = len(mKey)

const MaxLine = int32(^uint32(0) >> 1)

func ReadExcel() {
	f, err := excelize.OpenFile(FilePath)
	if err != nil {
		fmt.Println("读取excel错误:", err.Error())
		return
	}

	rows, err := f.Rows(SheetName)
	if err != nil {
		fmt.Println("遍历文件错误")
		return
	}

	mList := make([]map[string]string, 0)
	interval := int32(1)
	for rows.Next() {
		if interval == 1 {
			interval++
			continue
		}

		if interval >= MaxLine {
			break
		}

		row, _ := rows.Columns()
		line := make(map[string]string)

		for index, colCell := range row {
			if index >= mKeyLen {
				break
			}

			line[mKey[index]] = colCell
		}

		mList = append(mList, line)
		interval++
	}

	fmt.Println("导出数量：", len(mList))

	bString, err := json.Marshal(mList)
	if err != nil {
		fmt.Println("json序列化失败")
		return
	}

	err = ioutil.WriteFile(OutFilePath, bString, 0777)
	if err != nil {
		fmt.Println("生成文件失败")
		return
	}

	fmt.Println("生成文件成功")
}
