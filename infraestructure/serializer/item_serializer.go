package serializer

import "warehouse/domain"

type ItemSerializer struct {
	Id 						string
	Name                 	string
	UnitSizePresentation 	string
	SizePresentation     	int
	Code                 	string
	Container 				string
	Photo 					string
}

func NewItemSerializer(id string, name string, unitSizePresentation string, sizePresentation int,
	code string, container string, photo string) *ItemSerializer {
	return &ItemSerializer{Id: id, Name: name, UnitSizePresentation: unitSizePresentation,
		SizePresentation: sizePresentation, Code: code, Container: container, Photo: photo}
}

func NewItemSerializerFromDomain(item *domain.Item) *ItemSerializer {
	return NewItemSerializer(item.Id, item.Name, item.UnitSizePresentation,item.SizePresentation,
		item.Code, item.Container, item.Photo)
}

func (ise ItemSerializer) ToDomain() *domain.Item {
	return domain.NewItem(ise.Id, ise.Name, ise.UnitSizePresentation,ise.SizePresentation,
		ise.Code, ise.Container, ise.Photo)
}