package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

func main() {
	_reflectValue()
	//_reflectType()
	//_t1()
}

type User struct {
	Id   int    `json:"id" db:"name"`
	Name string `json:"name" db:"name"`
}

func (user User) GetName() string {
	return user.Name
}

func (user *User) SetName() {
}

func _t1() {
	// 1.注意：go语言中的true和false是非类型化的值，没有类型，也就无法计算具体的占用字节数
	//t := reflect.TypeOf(true)
	//log.Printf("t~Bits:%+v\n", t.Bits())

	// 可以是nil
	//t1 := reflect.TypeOf(nil)
	//log.Printf("t1:%+v\n", t1)

	//t := reflect.TypeOf(true)
	//log.Printf("t:%+v, 种类:%+v, 名称:%+v,\n", t, t.Kind(), t.Name())

	// typeof结果调用该方法要注意，可能panic
	//tElem := t.Elem()
	//log.Printf("tElem:%+v, 种类:%+v, 名称:%+v,\n", tElem, tElem.Kind(), tElem.Name())

	// 结构体的方法分为2部分，一种是带*的，一种是不带*的，如果使用指针反射以后调用，只能获取到带*的
	// 反之只能获取不带*的
}

func _reflectType() {
	user := &User{Id: 1, Name: "许磊"}
	// 反射类型的原信息，可以推断出具体的类型，源码中可以看到
	t := reflect.TypeOf(user)

	// 指针类型的话，是取不到名字的
	//log.Printf("t:%+v, 种类:%+v, 名称:%+v,\n", t, t.Kind(), t.Name())

	// t.Elem() 用来获取指针指向的元素类型，并且只能获取符合类型数据，例如array,map.chan,slice,ptr等
	tElem := t.Elem()
	//log.Printf("tElem:%+v, 种类:%+v, 名称:%+v,\n", tElem, tElem.Kind(), tElem.Name())

	// 获取结构体的属性
	for i := 0; i < tElem.NumField(); i++ {
		structField := tElem.Field(i)
		log.Printf("t~Field():%+v\n", structField)
		log.Printf("t~Field():%+v\n", structField.Tag.Get("json"))
	}

}

func _reflectValue() {
	user := &User{Id: 1, Name: "许磊"}
	// Value interface内部嵌入了rtype,所以可以使用rtype的方法
	v := reflect.ValueOf(user)
	//log.Printf("v:%+v, 种类:%+v, String():%+v,\n", v, v.Kind(), v.String())
	log.Printf("Kind:%+v, isNil:%+v\n", v.Kind(), v.IsNil())
	//vElem := v.Elem()
	//log.Printf("vElem:%+v, 种类:%+v, 名称:%+v,\n", vElem, vElem.Kind(), vElem.String())

	user1, ok := v.Interface().(User)
	log.Printf("user1:%+v, bool=%t\n", user1, ok)
}

// reprOfValue 所有的数据类型toString
func reprOfValue(val reflect.Value) string {
	switch vt := val.Interface().(type) {
	case bool:
		return strconv.FormatBool(vt)
	case error:
		return vt.Error()
	case float32:
		return strconv.FormatFloat(float64(vt), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(vt, 'f', -1, 64)
	case fmt.Stringer:
		return vt.String()
	case int:
		return strconv.Itoa(vt)
	case int8:
		return strconv.Itoa(int(vt))
	case int16:
		return strconv.Itoa(int(vt))
	case int32:
		return strconv.Itoa(int(vt))
	case int64:
		return strconv.FormatInt(vt, 10)
	case string:
		return vt
	case uint:
		return strconv.FormatUint(uint64(vt), 10)
	case uint8:
		return strconv.FormatUint(uint64(vt), 10)
	case uint16:
		return strconv.FormatUint(uint64(vt), 10)
	case uint32:
		return strconv.FormatUint(uint64(vt), 10)
	case uint64:
		return strconv.FormatUint(vt, 10)
	case []byte:
		return string(vt)
	default:
		return fmt.Sprint(val.Interface())
	}
}

func Repr(v interface{}) string {
	if v == nil {
		return ""
	}

	// if func (v *Type) String() string, we can't use Elem()
	switch vt := v.(type) {
	case fmt.Stringer:
		return vt.String()
	}

	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr && !val.IsNil() {
		val = val.Elem()
	}

	return reprOfValue(val)
}
