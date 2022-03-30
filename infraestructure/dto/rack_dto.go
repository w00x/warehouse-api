package dto

import "warehouse/domain"

type RackDto struct {
	Id		uint 	`json:"id" uri:"id"`
	Name 	string	`json:"name"`
	Code 	string	`json:"code"`
}

func NewRackDto(id uint, name string, code string) *RackDto {
	return &RackDto{Id: id, Name: name, Code: code}
}

func NewRackDtoFromDomain(rack *domain.Rack) *RackDto {
	if rack == nil {
		return nil
	}
	return NewRackDto(rack.Id, rack.Name, rack.Code)
}

func NewRackListDtoFromDomains(racks *[]domain.Rack) []*RackDto {
	var rackDtos []*RackDto
	for _, rack := range *racks {
		rackDtos = append(rackDtos,
			NewRackDtoFromDomain(&rack))
	}
	return rackDtos
}

func (r RackDto) ToDomain() *domain.Rack {
	return domain.NewRack(r.Id, r.Name, r.Code)
}