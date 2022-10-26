package domain

import "github.com/google/uuid"

type Rack struct {
	id   string
	Name string
	Code string
}

func NewRack(id string, name string, code string) *Rack {
	if id == "" {
		id = uuid.New().String()
	}

	return &Rack{id: id, Name: name, Code: code}
}

func (i Rack) Id() string {
	return i.id
}
