package application

import (
	"warehouse/domain"
	"warehouse/domain/repository"
	"warehouse/infrastructure/errors"
)

type RackApplication struct {
	rackRepository repository.IRackRepository
}

func NewRackApplication(rackRepository repository.IRackRepository) *RackApplication {
	return &RackApplication{rackRepository}
}

func (rackApplication *RackApplication) All() (*[]domain.Rack, errors.IBaseError) {
	return rackApplication.rackRepository.All()
}

func (rackApplication *RackApplication) Show(id string) (*domain.Rack, errors.IBaseError) {
	return rackApplication.rackRepository.Find(id)
}

func (rackApplication *RackApplication) Update(rack *domain.Rack) (*domain.Rack, errors.IBaseError) {
	return rackApplication.rackRepository.Update(rack)
}

func (rackApplication *RackApplication) Create(rack *domain.Rack) (*domain.Rack, errors.IBaseError) {
	return rackApplication.rackRepository.Create(rack)
}

func (rackApplication *RackApplication) Delete(rack *domain.Rack) errors.IBaseError {
	return rackApplication.rackRepository.Delete(rack)
}
