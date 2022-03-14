package models

import (
	"gorm.io/gorm"
	"warehouse/domain"
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

func FromRackDomainToModel(i *domain.Rack) *Rack {
	if i == nil {
		return nil
	}
	return NewRack(i.Id, i.Name, i.Code)
}

func (i Rack) ToDomain() *domain.Rack {
	return domain.NewRack(i.Id, i.Name, i.Code)
}