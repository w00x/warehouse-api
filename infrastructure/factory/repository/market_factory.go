package repository

import (
	"warehouse/domain/repository"
	"warehouse/infrastructure/errors"
	"warehouse/infrastructure/repository/gorm"
)

func MarketFactory(adapter string) (repository.IMarketRepository, errors.IBaseError) {
	if adapter == "gorm" {
		return gorm.NewMarketRepository(), nil
	}

	return nil, errors.NewNotFoundError("Todo Factory not found")
}
