package application

import (
	"warehouse/domain"
	"warehouse/domain/repository"
)

type ItemApplication struct {
	itemRepository repository.IItemRepository
}

func NewItemApplication(itemRepository repository.IItemRepository) *ItemApplication {
	return &ItemApplication{ itemRepository }
}

func (itemApplication *ItemApplication) All() ([]*domain.Item, error) {
	return itemApplication.itemRepository.All()
}

func (itemApplication *ItemApplication) Show(id string) (*domain.Item, error) {
	return itemApplication.itemRepository.Find(id)
}

func (itemApplication *ItemApplication) Update(id string, name string, unitSizePresentation string,
	sizePresentation int,
	code string, container string, photo string) error {
	return itemApplication.itemRepository.Update(id, name, unitSizePresentation, sizePresentation,
		code, container, photo)
}

func (itemApplication *ItemApplication) Create(name string, unitSizePresentation string,
	sizePresentation int, code string,
	container string, photo string) (*domain.Item, error) {
	return itemApplication.itemRepository.Create(name, unitSizePresentation, sizePresentation,
		code, container, photo)
}

func (itemApplication *ItemApplication) Delete(id string) error {
	return itemApplication.itemRepository.Delete(id)
}