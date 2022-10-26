package repository

import (
	"warehouse/domain/repository"
	"warehouse/infrastructure/errors"
	"warehouse/infrastructure/repository/gorm"
)

func InventoryFactory(adapter string) (repository.IInventoryRepository, errors.IBaseError) {
	if adapter == "gorm" {
		return gorm.NewInventoryRepository(), nil
	}

	return nil, errors.NewNotFoundError("Todo Factory not found")
}
