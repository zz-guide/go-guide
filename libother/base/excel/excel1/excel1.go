package util

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
)

func TestExcel() {
	f := excelize.NewFile()

	index := f.NewSheet("Sheet1")
	_ = f.SetCellValue("Sheet1", "A1", "姓名")
	_ = f.SetCellValue("Sheet1", "B1", "年龄")

	_ = f.SetCellValue("Sheet1", "A2", "许磊")
	_ = f.SetCellValue("Sheet1", "B2", "21")

	f.SetActiveSheet(index)
	err := f.SaveAs("./student.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

func ReadDataFromExcel(filePath string) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	columnStartStr := "A"
	currentColumnStr := columnStartStr
	columnStart := []rune(columnStartStr)
	currentColumn := []rune(currentColumnStr)

	key := ""
	columnEnd := []rune("Z")

	var rrr []map[string]string

	rows, err := f.GetRows("orders")
	for index, row := range rows {
		if index <= 1 {
			continue
		}

		data := make(map[string]string)
		key = strconv.Itoa(index)
		currentColumn[0] = columnStart[0]
		currentColumnStr = string(currentColumn)
		for _, item := range row {
			if currentColumn[0] >= columnEnd[0] {
				break
			}

			data[currentColumnStr+key] = item
			currentColumnStr = string(currentColumn[0] + 1)
			currentColumn[0] += 1
		}

		rrr = append(rrr, data)
	}

	//fmt.Println(rrr[0]["A2"])
	//fmt.Println(rrr[1])
	//fmt.Println(rrr[2])
	//fmt.Println(rrr[3])
}
