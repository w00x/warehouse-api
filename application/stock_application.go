package application

import (
	"warehouse/domain"
	"warehouse/domain/repository"
	"warehouse/infraestructure/errors"
)

type StockApplication struct {
	stockRepository repository.IStockRepository
}

func NewStockApplication(stockRepository repository.IStockRepository) *StockApplication {
	return &StockApplication{ stockRepository }
}

func (stockApplication *StockApplication) All() (*[]domain.Stock, errors.IBaseError) {
	return stockApplication.stockRepository.All()
}

func (stockApplication *StockApplication) Show(id uint) (*domain.Stock, errors.IBaseError) {
	return stockApplication.stockRepository.Find(id)
}

func (stockApplication *StockApplication) Update(stock *domain.Stock) (*domain.Stock, errors.IBaseError) {
	return stockApplication.stockRepository.Update(stock)
}

func (stockApplication *StockApplication) Create(stock *domain.Stock) (*domain.Stock, errors.IBaseError) {
	return stockApplication.stockRepository.Create(stock)
}

func (stockApplication *StockApplication) Delete(stock *domain.Stock) errors.IBaseError {
	return stockApplication.stockRepository.Delete(stock)
}