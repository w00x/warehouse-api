package models

import (
	"gorm.io/gorm"
)

type Market struct {
	gorm.Model
	Id			uint		`gorm:"primaryKey;autoIncrement"`
	Name 		string
	Prices		[]Price
}

func NewMarket(id uint, name string) *Market {
	return &Market{Id: id, Name: name}
}