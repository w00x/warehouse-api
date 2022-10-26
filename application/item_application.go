package application

import (
	"warehouse/domain"
	"warehouse/domain/repository"
	"warehouse/infrastructure/errors"
)

type ItemApplication struct {
	itemRepository repository.IItemRepository
}

func NewItemApplication(itemRepository repository.IItemRepository) *ItemApplication {
	return &ItemApplication{itemRepository}
}

func (itemApplication *ItemApplication) All() (*[]domain.Item, errors.IBaseError) {
	return itemApplication.itemRepository.All()
}

func (itemApplication *ItemApplication) Show(id string) (*domain.Item, errors.IBaseError) {
	return itemApplication.itemRepository.Find(id)
}

func (itemApplication *ItemApplication) Update(item *domain.Item) (*domain.Item, errors.IBaseError) {
	return itemApplication.itemRepository.Update(item)
}

func (itemApplication *ItemApplication) Create(item *domain.Item) (*domain.Item, errors.IBaseError) {
	return itemApplication.itemRepository.Create(item)
}

func (itemApplication *ItemApplication) Delete(item *domain.Item) errors.IBaseError {
	return itemApplication.itemRepository.Delete(item)
}
