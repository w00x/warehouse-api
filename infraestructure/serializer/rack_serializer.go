package serializer

import "warehouse/domain"

type RackSerializer struct {
	Id		uint 	`json:"id" uri:"id"`
	Name 	string	`json:"name"`
	Code 	string	`json:"code"`
}

func NewRackSerializer(id uint, name string, code string) *RackSerializer {
	return &RackSerializer{Id: id, Name: name, Code: code}
}

func NewRackSerializerFromDomain(rack *domain.Rack) *RackSerializer {
	if rack == nil {
		return nil
	}
	return NewRackSerializer(rack.Id, rack.Name, rack.Code)
}

func NewRackListSerializerFromDomains(racks *[]domain.Rack) []*RackSerializer {
	var rackSerializers []*RackSerializer
	for _, rack := range *racks {
		rackSerializers = append(rackSerializers,
			NewRackSerializerFromDomain(&rack))
	}
	return rackSerializers
}

func (r RackSerializer) ToDomain() *domain.Rack {
	return domain.NewRack(r.Id, r.Name, r.Code)
}