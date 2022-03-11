package models

import (
	"gorm.io/gorm"
	"warehouse/domain"
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

func FromMarketDomainToModel(i *domain.Market) *Market {
	if i == nil {
		return nil
	}
	return NewMarket(i.Id, i.Name)
}

func (i Market) ToDomain() *domain.Market {
	return domain.NewMarket(i.Id, i.Name)
}