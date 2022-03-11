package serializer

import "warehouse/domain"

type ItemSerializer struct {
	Id 						uint 	`json:"id" uri:"id"`
	Name                 	string 	`json:"name"`
	UnitSizePresentation 	string 	`json:"unit_size_presentation"`
	SizePresentation     	int 	`json:"size_presentation"`
	Code                 	string 	`json:"code"`
	Container 				string 	`json:"container"`
	Photo 					string 	`json:"photo"`
}

func NewItemSerializer(id uint, name string, unitSizePresentation string, sizePresentation int,
	code string, container string, photo string) *ItemSerializer {
	return &ItemSerializer{Id: id, Name: name, UnitSizePresentation: unitSizePresentation,
		SizePresentation: sizePresentation, Code: code, Container: container, Photo: photo}
}

func NewItemSerializerFromDomain(item *domain.Item) *ItemSerializer {
	if item == nil {
		return nil
	}
	return NewItemSerializer(item.Id, item.Name, item.UnitSizePresentation,item.SizePresentation,
		item.Code, item.Container, item.Photo)
}

func NewItemListSerializerFromDomains(items *[]domain.Item) []*ItemSerializer {
	var itemSerializers []*ItemSerializer
	for _, item := range *items {
		itemSerializers = append(itemSerializers,
			NewItemSerializerFromDomain(&item))
	}
	return itemSerializers
}

func (ise ItemSerializer) ToDomain() *domain.Item {
	return domain.NewItem(ise.Id, ise.Name, ise.UnitSizePresentation,ise.SizePresentation,
		ise.Code, ise.Container, ise.Photo)
}