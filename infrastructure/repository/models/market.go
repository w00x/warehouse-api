package models

import (
	"gorm.io/gorm"
)

type Market struct {
	gorm.Model
	ID     string `gorm:"type:uuid;primaryKey;column:id"`
	Name   string
	Prices []Price
}

func NewMarket(id string, name string) *Market {
	return &Market{ID: id, Name: name}
}
