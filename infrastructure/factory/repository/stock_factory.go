package repository

import (
	"warehouse/domain/repository"
	"warehouse/infrastructure/errors"
	"warehouse/infrastructure/repository/gorm"
)

func StockFactory(adapter string) (repository.IStockRepository, errors.IBaseError) {
	if adapter == "gorm" {
		return gorm.NewStockRepository(), nil
	}

	return nil, errors.NewNotFoundError("Todo Factory not found")
}
