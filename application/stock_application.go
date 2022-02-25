package application

import (
	"time"
	"warehouse/domain"
	"warehouse/domain/repository"
)

type StockApplication struct {
	stockRepository repository.IStockRepository
}

func NewStockApplication(stockRepository repository.IStockRepository) *StockApplication {
	return &StockApplication{ stockRepository }
}

func (stockApplication *StockApplication) All() ([]*domain.Stock, error) {
	return stockApplication.stockRepository.All()
}

func (stockApplication *StockApplication) Show(id string) (*domain.Stock, error) {
	return stockApplication.stockRepository.Find(id)
}

func (stockApplication *StockApplication) Update(id string, item *domain.Item, rack *domain.Rack,
	quantity int, operationDate time.Time, expirationDate time.Time) error {
	return stockApplication.stockRepository.Update(id, item, rack, quantity, operationDate,
		expirationDate)
}

func (stockApplication *StockApplication) Create(item *domain.Item, rack *domain.Rack,
	quantity int, operationDate time.Time,
	expirationDate time.Time) (*domain.Stock, error) {
	return stockApplication.stockRepository.Create(item, rack, quantity, operationDate,
		expirationDate)
}

func (stockApplication *StockApplication) Delete(id string) error {
	return stockApplication.stockRepository.Delete(id)
}