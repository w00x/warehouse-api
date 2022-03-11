package domain

type Item struct {
	Id 						uint
	Name                 	string
	UnitSizePresentation 	string
	SizePresentation     	int
	Code                 	string
	Container 				string
	Photo 					string
}

func NewItem(id uint, name string, unitSizePresentation string, sizePresentation int,
	code string, container string, photo string) *Item {
	return &Item{Id: id, Name: name, UnitSizePresentation: unitSizePresentation,
		SizePresentation: sizePresentation, Code: code, Container: container, Photo: photo}
}