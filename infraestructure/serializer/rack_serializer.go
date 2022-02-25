package serializer

import "warehouse/domain"

type RackSerializer struct {
	Id		string
	Name 	string
	Code 	string
}

func NewRackSerializer(id string, name string, code string) *RackSerializer {
	return &RackSerializer{Id: id, Name: name, Code: code}
}

func NewRackSerializerFromDomain(rack *domain.Rack) *RackSerializer {
	return NewRackSerializer(rack.Id, rack.Name, rack.Code)
}

func (r RackSerializer) ToDomain() *domain.Rack {
	return domain.NewRack(r.Id, r.Name, r.Code)
}