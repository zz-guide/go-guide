package model

type Car struct {
	Model
	Name      string
	StudentId int64
	SID       int64
}

func (Car) TableName() string {
	return "car"
}
