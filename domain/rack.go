package domain

type Rack struct {
	Id		string
	Name 	string
	Code 	string
}

func NewRack(id string, name string, code string) *Rack {
	return &Rack{Id: id, Name: name, Code: code}
}
