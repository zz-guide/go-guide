package Json

import (
	"encoding/json"
	"fmt"
)

/**
结论：
①某个字段为空是不会映射到结构体中的
②字段的类型要一一对应
③可以起别名映射json字段
④在预先知道json结构的时候可以直接映射到结构体
⑤不知道json结构的时候可以映射到空Interface
⑥要实现转化，字段首字母必须是大写的
⑦
*/
func getJson() string {
	return `{
    "errcode": 200,
    "errmsg": "success",
    "errtime": "1524988100",
    "data": {
        "student_unit_id":"1",
        "unit_no": "3",
        "button_time": {
            "composite_spell_stay_time": "1500",
            "whole_spell_stay_time": "1000"
        },
        "list": [
            {
                "type": "1",
                "words": [
                    {
                        "word_id": "5932",
                        "word": "she"
                    },
                    {
                        "word_id": "6830",
                        "word": "too"
                    }
                ]
            },
            {
                "type": "2",
                "words": [
                    {
                        "word_id": "2771",
                        "word": "friend"
                    },
                    {
                        "word_id": "6029",
                        "word": "sister"
                    }
                ]
            }
        ]
    }
}`
}

/**
全部数据结点
`json:"errcode,omitempty"` ,表示为空则不输出到json串
*/
type CheckList struct {
	ID      int `json:"-"`
	Errcode int `json:"errcode,omitempty"`
	Errmsg  string
	Errtime string
	Data    DataSet
}

/**
data结点
*/
type DataSet struct {
	Student_unit_id string
	Unit_no         string
	Button_time     ButtonTime
	List            []WordList
}

/**
button_time结点
*/
type ButtonTime struct {
	Composite_spell_stay_time string
	Whole_spell_stay_time     string
}

/**
  json  tag 用来映射字段，比如正好json中有go的关键字，可以起一个别名

*/
type WordList struct {
	MyType string `json:"type"`
	Words  []Words
}

type Words struct {
	Word_id string
	Word    string
}

func Test1() {
	var s CheckList
	jsonStr := getJson()
	json.Unmarshal([]byte(jsonStr), &s)

	fmt.Println(s)
}
