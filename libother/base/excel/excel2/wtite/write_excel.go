/**
作用：用于将json文件中的内容按照配置生成excel文件。
*/
package wtite

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

const OutFilePath = "/Users/xulei/Downloads/生成Excel表格.xlsx" //生成excel文件的绝对路径
const InFilePath = "/Users/xulei/Downloads/读取json.txt"      //读取json文件的绝对路径
const SheetName = "Sheet1"                                  //默认只支持第一个sheet

//字段配置，excel索引行，中文title,json中对应的字段
var excelConfig = [][]string{
	{"A", "学段", "stage"},
	{"B", "教材版本", "material"},
	{"C", "课本", "book"},
	{"D", "有效期", "valid"},
}

var excelCategories = make(map[string]string, 0)
var excelKeys = make([]string, 0)
var excelLines = make([]string, 0)

const StartIndex = 2

type ColumnData struct {
	Line   string
	Column string
	Key    string
}

func init() {
	for _, item := range excelConfig {
		excelCategories[item[0]] = item[1]
		excelKeys = append(excelKeys, item[2])
		excelLines = append(excelLines, item[0])
	}
}

func CreateExcel() {
	values, err := _getJsonContent()
	if err != nil {
		fmt.Println("获取json内容失败:", err)
		return
	}

	f := excelize.NewFile()
	//设置表头内容
	for k, v := range excelCategories {
		f.SetCellValue(SheetName, k+"1", v)
	}

	startIndex := StartIndex
	//设置行
	for _, item := range values {
		var build strings.Builder
		for index, v := range excelKeys {
			_, ok := item[v]
			if !ok {
				continue
			}

			build.WriteString(excelLines[index])
			build.WriteString(strconv.Itoa(startIndex))
			f.SetCellValue(SheetName, build.String(), item[v])
			build.Reset()
		}

		startIndex++
	}

	if err := f.SaveAs(OutFilePath); err != nil {
		fmt.Println("生成excel文件失败:", err)
		return
	}

	fmt.Println("生成excel文件成功")
}

func _getJsonContent() ([]map[string]interface{}, error) {
	var mapResult []map[string]interface{}
	var err error = nil

	jsonByte, err := ioutil.ReadFile(InFilePath)
	if err != nil {
		fmt.Println("读取json文件失败")
		return mapResult, err
	}

	//TODO::本地测试可以放开这2行代码
	//jsonStr := `[{"book":"小学考纲词库","material":"考纲版","stage":"小学","valid":"一年"},{"book":"小学高分培优词库","material":"考纲版","stage":"小学","valid":"一年"}]`
	//jsonByte := []byte(jsonStr)

	err = json.Unmarshal(jsonByte, &mapResult)
	if err != nil {
		return mapResult, err
	}

	return mapResult, err
}
