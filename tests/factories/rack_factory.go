package factories

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"warehouse/domain"
	"warehouse/infrastructure/repository/gorm"
)

type Rack struct {
	Name string
	Code string
}

func (i Rack) ToDomain() *domain.Rack {
	return domain.NewRack("", i.Name, i.Code)
}

func FromRackDomainToFactory(rack *domain.Rack) *Rack {
	return &Rack{
		Name: rack.Name,
		Code: rack.Code,
	}
}

func NewRackFactory() *domain.Rack {
	Rack := &Rack{}
	err := gofakeit.Struct(Rack)
	if err != nil {
		fmt.Println(err)
	}

	repo := gorm.NewRackRepository()
	RackDomain, errRepo := repo.Create(Rack.ToDomain())
	if errRepo != nil {
		panic(err)
	}

	return RackDomain
}

func NewRackDomainFactory() *domain.Rack {
	Rack := &Rack{}
	err := gofakeit.Struct(Rack)
	if err != nil {
		fmt.Println(err)
	}

	return Rack.ToDomain()
}

func NewRackObjectFactory() map[string]interface{} {
	rack := &Rack{}
	err := gofakeit.Struct(rack)
	if err != nil {
		fmt.Println(err)
	}

	rackMarshal := map[string]interface{}{
		"name": rack.Name,
		"code": rack.Code,
	}

	return rackMarshal
}

func NewRackFactoryList(count int) []*domain.Rack {
	var RackDomains []*domain.Rack
	repo := gorm.NewRackRepository()

	for i := 0; i < count; i++ {
		rack := &Rack{}
		err := gofakeit.Struct(rack)
		if err != nil {
			panic(err)
		}

		RackDomain, errRepo := repo.Create(rack.ToDomain())
		if errRepo != nil {
			panic(err)
		}
		RackDomains = append(RackDomains, RackDomain)
	}

	return RackDomains
}
