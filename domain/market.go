package domain

import "github.com/google/uuid"

type Market struct {
	id   string
	Name string
}

func NewMarket(id string, name string) *Market {
	if id == "" {
		id = uuid.New().String()
	}

	return &Market{id: id, Name: name}
}

func (i Market) Id() string {
	return i.id
}
