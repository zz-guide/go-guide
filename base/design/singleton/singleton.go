package singleton

import (
	"sync"
)

var (
	stu  *Student
	once sync.Once
)

type Student struct {
	Name string
}

func GetInstance() *Student {
	once.Do(func() {
		stu = &Student{Name: "许磊"}
	})

	return stu
}
