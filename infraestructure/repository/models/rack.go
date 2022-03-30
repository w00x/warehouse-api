package models

import (
	"gorm.io/gorm"
)

type Rack struct {
	gorm.Model
	Id			uint		`gorm:"primaryKey;autoIncrement"`
	Name 		string
	Code 		string
	Stocks		[]Stock
}

func NewRack(id uint, name string, code string) *Rack {
	return &Rack{Id: id, Name: name, Code: code}
}