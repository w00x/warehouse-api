package repository

import (
	"warehouse/domain/repository"
	"warehouse/infrastructure/errors"
	"warehouse/infrastructure/repository/gorm"
)

func ItemFactory(adapter string) (repository.IItemRepository, errors.IBaseError) {
	if adapter == "gorm" {
		return gorm.NewItemRepository(), nil
	}

	return nil, errors.NewNotFoundError("Todo Factory not found")
}
