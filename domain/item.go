package domain

import "github.com/google/uuid"

type Item struct {
	id                   string
	Name                 string
	UnitSizePresentation string
	SizePresentation     int
	Code                 string
	Container            string
	Photo                string
}

func NewItem(id string, name string, unitSizePresentation string, sizePresentation int,
	code string, container string, photo string) *Item {
	if id == "" {
		id = uuid.New().String()
	}

	return &Item{id: id, Name: name, UnitSizePresentation: unitSizePresentation,
		SizePresentation: sizePresentation, Code: code, Container: container, Photo: photo}
}

func (i Item) Id() string {
	return i.id
}
