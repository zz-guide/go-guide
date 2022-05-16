package main

import (
	"database/sql"
	"github.com/tal-tech/go-zero/core/mapping"
	"log"
	"reflect"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func init() {
	db, err = sql.Open("mysql", "root:xl123456?@tcp(47.105.50.31:3306)/casbin?charset=utf8mb4&parseTime=True")
	if err != nil {
		log.Println("open mysql failed,", err)
		return
	}

	db.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
	db.SetMaxOpenConns(100)                  //设置最大连接数
	db.SetMaxIdleConns(16)                   //设置闲置连接数

	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		log.Println("ping mysql failed,", err)
		return
	}

	//defer func() {
	//	_ = database.Close()
	//}()
}

type User struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

/**

 */

func main() {
	_query()
}

func Deref(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t
}

func unwrapFields(v reflect.Value) []reflect.Value {
	var fields []reflect.Value
	// 校验ptr和struct
	indirect := reflect.Indirect(v)
	if indirect.Kind() != reflect.Struct {
		panic("indirect must be Struct Kind!")
	}

	log.Println("indirect:", indirect.Kind())

	for i := 0; i < indirect.NumField(); i++ {
		child := indirect.Field(i)
		// 结构体属性可能是指针
		if child.Kind() == reflect.Ptr && child.IsNil() {
			baseValueType := Deref(child.Type())
			// 因为reflect.New返回的是指向baseValueType的指针，如果baseValueType也是指针的话，相当于是指针的指针，所以
			// 需要baseValueType是值类型才可以
			child.Set(reflect.New(baseValueType))
		}

		child = reflect.Indirect(child)
		childType := indirect.Type().Field(i)
		if child.Kind() == reflect.Struct && childType.Anonymous {
			fields = append(fields, unwrapFields(child)...)
		} else {
			fields = append(fields, child)
		}
	}

	return fields
}

func _query() {
	sqlStr := "select `id`,`name` from `user`"
	rows, err := db.Query(sqlStr)
	if err != nil {
		log.Println("查询失败, ", err, rows)
		return
	}

	defer rows.Close()

	// Scan分几种情况
	// 1.struct，scanner.Columns()获取所有的列，然后反射组装
	// 2.基本数据类型，基本变量，
	columns, err := rows.Columns()
	if err != nil {
		log.Println("Columns 失败", err)
		return
	} else {
		log.Println("columns:", columns, len(columns))
	}

	for rows.Next() {
		// 此时是nil，还需要初始化才能使用
		var v User
		// 变量地址反射
		rv := reflect.ValueOf(&v)
		// Value从ptr->Struct的转变
		fields := unwrapFields(rv.Elem())
		values := make([]interface{}, len(columns))
		for i := 0; i < len(values); i++ {
			valueField := fields[i]
			switch valueField.Kind() {
			case reflect.Ptr:
				if !valueField.CanInterface() {
					return
				}
				if valueField.IsNil() {
					baseValueType := mapping.Deref(valueField.Type())
					valueField.Set(reflect.New(baseValueType))
				}
				values[i] = valueField.Interface()
			default:
				if !valueField.CanAddr() || !valueField.Addr().CanInterface() {
					return
				}
				values[i] = valueField.Addr().Interface()
			}
		}

		log.Println("values:", len(values))
		if err := rows.Scan(values...); err != nil {
			log.Println("解析失败, ", err, rows)
			return
		}

		log.Printf("user=%+v\n", v)
	}
}
