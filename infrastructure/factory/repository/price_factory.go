package repository

import (
	"warehouse/domain/repository"
	"warehouse/infrastructure/errors"
	"warehouse/infrastructure/repository/gorm"
)

func PriceFactory(adapter string) (repository.IPriceRepository, errors.IBaseError) {
	if adapter == "gorm" {
		return gorm.NewPriceRepository(), nil
	}

	return nil, errors.NewNotFoundError("Todo Factory not found")
}
