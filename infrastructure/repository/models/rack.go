package models

import (
	"gorm.io/gorm"
)

type Rack struct {
	gorm.Model
	ID     string `gorm:"type:uuid;primaryKey;column:id"`
	Name   string
	Code   string
	Stocks []Stock
}

func NewRack(id string, name string, code string) *Rack {
	return &Rack{ID: id, Name: name, Code: code}
}
