package domain

type Market struct {
	Id		string
	Name 	string
}

func NewMarket(id string, name string) *Market {
	return &Market{Id: id, Name: name}
}
