package domain

type Market struct {
	Id			uint
	Name 		string
}

func NewMarket(id uint, name string) *Market {
	return &Market{Id: id, Name: name}
}
