package factories

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"testing"
	"warehouse/domain"
	"warehouse/infraestructure/repository/postgres"
)

type Rack struct {
	Name 	string
	Code 	string
}

func (i Rack) ToDomain() *domain.Rack {
	return domain.NewRack(0, i.Name, i.Code)
}

func FromRackDomainToFactory(rack *domain.Rack) *Rack {
	return &Rack{
		Name: rack.Name,
		Code: rack.Code,
	}
}

func NewRackFactory(t *testing.T) *domain.Rack {
	Rack := &Rack{}
	err := faker.FakeData(Rack)
	if err != nil {
		fmt.Println(err)
	}

	repo := postgres.NewRackRepository()
	RackDomain, errRepo := repo.Create(Rack.ToDomain())
	if errRepo != nil {
		panic(err)
	}

	t.Cleanup(func() {
		CleanRack()
	})

	return RackDomain
}

func NewRackObjectFactory() map[string]interface{} {
	rack := &Rack{}
	err := faker.FakeData(rack)
	if err != nil {
		fmt.Println(err)
	}

	rackMarshal := map[string]interface{}{
		"name": rack.Name,
		"code": rack.Code,
	}

	return rackMarshal
}

func NewRackFactoryList(count int, t *testing.T) []*domain.Rack {
	var RackDomains []*domain.Rack
	repo := postgres.NewRackRepository()

	for i := 0; i < count; i++ {
		Rack := &Rack{}
		err := faker.FakeData(Rack)
		if err != nil {
			panic(err)
		}

		RackDomain, errRepo := repo.Create(Rack.ToDomain())
		if errRepo != nil {
			panic(err)
		}
		RackDomains = append(RackDomains, RackDomain)
	}

	t.Cleanup(func() {
		CleanRack()
	})

	return RackDomains
}

func CleanRack() {
	postgres.NewPostgresBase().DB.Exec("DELETE FROM racks")
	postgres.NewPostgresBase().DB.Exec("ALTER SEQUENCE racks_id_seq RESTART WITH 1")
}
