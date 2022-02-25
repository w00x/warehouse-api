package application

import (
	"warehouse/domain"
	"warehouse/domain/repository"
)

type RackApplication struct {
	rackRepository repository.IRackRepository
}

func NewRackApplication(rackRepository repository.IRackRepository) *RackApplication {
	return &RackApplication{ rackRepository }
}

func (rackApplication *RackApplication) All() ([]*domain.Rack, error) {
	return rackApplication.rackRepository.All()
}

func (rackApplication *RackApplication) Show(id string) (*domain.Rack, error) {
	return rackApplication.rackRepository.Find(id)
}

func (rackApplication *RackApplication) Update(id string, name string, code string) error {
	return rackApplication.rackRepository.Update(id, name, code)
}

func (rackApplication *RackApplication) Create(name string, code string) (*domain.Rack, error) {
	return rackApplication.rackRepository.Create(name, code)
}

func (rackApplication *RackApplication) Delete(id string) error {
	return rackApplication.rackRepository.Delete(id)
}