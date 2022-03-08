package postgres

import (
	"gorm.io/gorm"
	"warehouse/domain"
	"warehouse/infraestructure/errors"
)

type InventoryRepository struct {
	postgresBase *PostgresBase
}

func NewInventoryRepository() *InventoryRepository {
	postgresBase := NewPostgresBase()
	return &InventoryRepository{postgresBase}
}

func (i InventoryRepository) All() (*[]domain.Inventory, errors.IBaseError) {
	var instances []domain.Inventory
	result := i.postgresBase.DB.Model(&domain.Inventory{}).Scan(&instances)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &instances, nil
}

func (i InventoryRepository) Find(id uint) (*domain.Inventory, errors.IBaseError) {
	var instance domain.Inventory
	result := i.postgresBase.DB.First(&instance, id)

	if err := result.Error; err == gorm.ErrRecordNotFound {
		return nil, errors.NewNotFoundError("Repository not found")
	} else if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	return &instance, nil
}

func (i InventoryRepository) Create(instance *domain.Inventory) (*domain.Inventory, errors.IBaseError) {
	result:= i.postgresBase.DB.Create(instance)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Repository not created")
	}

	return instance, nil
}

func (i InventoryRepository) Update(instance *domain.Inventory) (*domain.Inventory, errors.IBaseError) {
	result := i.postgresBase.DB.Save(instance)

	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Repository not updated")
	}

	return instance, nil
}

func (i InventoryRepository) Delete(instance *domain.Inventory) errors.IBaseError {
	result := i.postgresBase.DB.Delete(instance)

	if err := result.Error; err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return errors.NewNotFoundError("Repository not deleted")
	}

	return nil
}