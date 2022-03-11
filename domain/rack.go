package domain

type Rack struct {
	Id			uint
	Name 		string
	Code 		string
}

func NewRack(id uint, name string, code string) *Rack {
	return &Rack{Id: id, Name: name, Code: code}
}
