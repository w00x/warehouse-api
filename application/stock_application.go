package application

import (
	"warehouse/domain"
	"warehouse/domain/repository"
	"warehouse/infrastructure/errors"
)

type StockApplication struct {
	stockRepository repository.IStockRepository
}

func NewStockApplication(stockRepository repository.IStockRepository) *StockApplication {
	return &StockApplication{stockRepository}
}

func (stockApplication *StockApplication) All() (*[]domain.Stock, errors.IBaseError) {
	return stockApplication.stockRepository.All()
}

func (stockApplication *StockApplication) Show(id string) (*domain.Stock, errors.IBaseError) {
	return stockApplication.stockRepository.Find(id)
}

func (stockApplication *StockApplication) Create(stock *domain.Stock) (*domain.Stock, errors.IBaseError) {
	return stockApplication.stockRepository.Create(stock)
}
