package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ID                   string `gorm:"type:uuid;primaryKey;column:id"`
	Name                 string
	UnitSizePresentation string
	SizePresentation     int
	Code                 string
	Container            string
	Photo                string
	Prices               []Price
	Stocks               []Stock
}

func NewItem(id string, name string, unitSizePresentation string, sizePresentation int,
	code string, container string, photo string) *Item {
	return &Item{ID: id, Name: name, UnitSizePresentation: unitSizePresentation,
		SizePresentation: sizePresentation, Code: code, Container: container, Photo: photo}
}
