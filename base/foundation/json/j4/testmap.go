package Json

import (
	"encoding/json"
	"fmt"
)

func getMapJson() string {
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

func TestMap() {
	b := []byte(getMapJson())
	var f interface{}
	json.Unmarshal(b, &f)
	fmt.Println(f)
	//这句是断言语句
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is mystring", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

}
