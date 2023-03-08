package application

import (
	"warehouse/domain"
	"warehouse/domain/repository"
	"warehouse/infrastructure/errors"
	"warehouse/shared"
)

type StockApplication struct {
	stockRepository repository.IStockRepository
	itemRepository  repository.IItemRepository
	rackRepository  repository.IRackRepository
}

func NewStockApplication(stockRepository repository.IStockRepository, itemRepository repository.IItemRepository, rackRepository repository.IRackRepository) *StockApplication {
	return &StockApplication{stockRepository, itemRepository, rackRepository}
}

func (stockApplication *StockApplication) All() (*[]domain.Stock, errors.IBaseError) {
	return stockApplication.stockRepository.AllByLastInventory()
}

func (stockApplication *StockApplication) Show(id string) (*domain.Stock, errors.IBaseError) {
	return stockApplication.stockRepository.Find(id)
}

func (stockApplication *StockApplication) Create(itemCode string, rackCode string, quantity int, operationDate shared.DateTime, expirationDate shared.DateTime, comment string) (*domain.Stock, errors.IBaseError) {
	item, errItem := stockApplication.itemRepository.FindByCode(itemCode)
	if errItem != nil {
		return nil, errItem
	}
	rack, errRack := stockApplication.rackRepository.FindByCode(rackCode)
	if errRack != nil {
		return nil, errRack
	}
	stock := domain.NewStock("", item, rack, quantity, operationDate, expirationDate, comment, item.Id(), rack.Id())
	return stockApplication.stockRepository.Create(stock)
}

func (stockApplication StockApplication) AllByInventory(inventoryId string) (*[]domain.Stock, errors.IBaseError) {
	return stockApplication.stockRepository.AllByInventory(inventoryId)
}
