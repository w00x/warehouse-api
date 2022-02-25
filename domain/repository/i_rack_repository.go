package repository

import "warehouse/domain"

type IRackRepository interface {
	All() 										([]*domain.Rack, error)
	Find(id string) 							(*domain.Rack, error)
	Update(id string, name string, code string) error
	Create(name string, code string) 			(*domain.Rack, error)
	Delete(id string) 							error
}
