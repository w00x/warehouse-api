package repository

import "warehouse/domain"

type IItemRepository interface {
	All() ([]*domain.Item, error)
	Find(id string) (*domain.Item, error)
	Update(id string, name string, unit_size_presentation string, size_presentation int,
		code string, container string, photo string) error
	Create(name string, unit_size_presentation string, size_presentation int, code string,
		container string, photo string) (*domain.Item, error)
	Delete(id string) error
}
