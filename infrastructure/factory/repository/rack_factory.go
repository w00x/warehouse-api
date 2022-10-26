package repository

import (
	"warehouse/domain/repository"
	"warehouse/infrastructure/errors"
	"warehouse/infrastructure/repository/gorm"
)

func RackFactory(adapter string) (repository.IRackRepository, errors.IBaseError) {
	if adapter == "gorm" {
		return gorm.NewRackRepository(), nil
	}

	return nil, errors.NewNotFoundError("Todo Factory not found")
}
