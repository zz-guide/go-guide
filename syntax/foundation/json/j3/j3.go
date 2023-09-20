package jjson

import (
	"encoding/json"
	"log"
)

type Person struct {
	//Name string `json:"name"` //TODO name
	Sex string `json:"sex"`
}

type Student struct {
	Name string `json:"name"` //TODO name
	Id   string `json:"id"`
}

type User struct {
	*Person
	*Student // TODO name会重复.
}

// 结论，2个结构体若有相同的属性且都是公开访问的，json会丢失该属性
func D() {
	str := `{
		  "name": "张三",
		  "sex" : "男",
		  "Id": "10001"
		}`

	user := new(User)
	err := json.Unmarshal([]byte(str), &user)
	log.Println("ERR:", err, user.Student)
	str1, err := json.Marshal(user)

	log.Println("user json==>", string(str1))
}
