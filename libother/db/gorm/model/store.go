package model

type Store struct {
	Model
	Name string
}

func (Store) TableName() string {
	return "store"
}
