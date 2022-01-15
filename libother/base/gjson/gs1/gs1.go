package bson

import (
	"fmt"
	"github.com/tidwall/gjson"
)

func TestJson() {
	const json = `{"name":{"first":"Janet","last1":null},"age":47}`

	value := gjson.Get(json, "name.last1")

	fmt.Println(value)
}
