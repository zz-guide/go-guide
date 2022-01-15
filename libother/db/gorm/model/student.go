package model

import (
	"time"
)

type Student struct {
	Model
	Name     string
	Age      uint
	StoreId  int64
	CarId    int64 `gorm:"index"`
	Car      Car   `gorm:"foreignKey:StudentId;references:CarId"`
	Store    Store
	Cars     []Car      `gorm:"foreignKey:SID;"` // Car表如果有关联ID的话，只需要配置foreignKey，相当于Has Many
	Birthday *time.Time `gorm:"default:0000-00-01"`
}

func (Student) TableName() string {
	return "student"
}
