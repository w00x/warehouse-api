package models

import (
	"gorm.io/gorm"
	"warehouse/domain"
)

type Item struct {
	gorm.Model
	Id 						uint		`gorm:"primaryKey;autoIncrement"`
	Name                 	string
	UnitSizePresentation 	string
	SizePresentation     	int
	Code                 	string
	Container 				string
	Photo 					string
	Prices					[]Price
	Stocks					[]Stock
}

func NewItem(id uint, name string, unitSizePresentation string, sizePresentation int,
	code string, container string, photo string) *Item {
	return &Item{Id: id, Name: name, UnitSizePresentation: unitSizePresentation,
		SizePresentation: sizePresentation, Code: code, Container: container, Photo: photo}
}

func FromItemDomainToModel(i *domain.Item) *Item {
	if i == nil {
		return nil
	}
	return NewItem(i.Id, i.Name, i.UnitSizePresentation, i.SizePresentation, i.Code, i.Container, i.Photo)
}

func (i Item) ToDomain() *domain.Item {
	return domain.NewItem(i.Id, i.Name, i.UnitSizePresentation, i.SizePresentation, i.Code, i.Container, i.Photo)
}