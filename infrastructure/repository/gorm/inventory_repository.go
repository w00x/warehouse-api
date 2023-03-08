package gorm

import (
	"gorm.io/gorm"
	"warehouse/domain"
	"warehouse/infrastructure/errors"
	"warehouse/infrastructure/repository/mappers"
	"warehouse/infrastructure/repository/models"
)

type InventoryRepository struct {
	postgresBase *PostgresBase
}

func NewInventoryRepository() *InventoryRepository {
	postgresBase := NewPostgresBase()
	return &InventoryRepository{postgresBase}
}

func (r InventoryRepository) All() (*[]domain.Inventory, errors.IBaseError) {
	var instances []models.Inventory
	result := r.postgresBase.DB.Model(&models.Inventory{}).Scan(&instances)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return mappers.NewInventoryListDomainFromModel(&instances), nil
}

func (r InventoryRepository) Find(id string) (*domain.Inventory, errors.IBaseError) {
	var instance models.Inventory
	result := r.postgresBase.DB.First(&instance, "id = ?", id)

	if err := result.Error; err == gorm.ErrRecordNotFound {
		return nil, errors.NewNotFoundError("Repository not found")
	} else if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	return mappers.FromInventoryModelToDomain(&instance), nil
}

func (r InventoryRepository) Create(instance *domain.Inventory) (*domain.Inventory, errors.IBaseError) {
	model := mappers.FromInventoryDomainToModel(instance)
	result := r.postgresBase.DB.Create(model)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Repository not created")
	}

	return r.Find(model.ID)
}

func (r InventoryRepository) Update(instance *domain.Inventory) (*domain.Inventory, errors.IBaseError) {
	instanceModel := mappers.FromInventoryDomainToModel(instance)
	d, err := r.Find(instanceModel.ID)
	if err != nil {
		return nil, err
	}

	model := mappers.FromInventoryDomainToModel(d)
	result := r.postgresBase.DB.Model(model).Updates(instanceModel)

	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Repository not updated")
	}

	return r.Find(instance.Id())
}

func (r InventoryRepository) Delete(instance *domain.Inventory) errors.IBaseError {
	model := mappers.FromInventoryDomainToModel(instance)
	result := r.postgresBase.DB.Delete(model)

	if err := result.Error; err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return errors.NewNotFoundError("Repository not deleted")
	}

	return nil
}

func (r InventoryRepository) Last() (*domain.Inventory, errors.IBaseError) {
	var instance models.Inventory
	result := r.postgresBase.DB.Order("operation_date DESC").First(&instance)

	if err := result.Error; err == gorm.ErrRecordNotFound {
		return nil, errors.NewNotFoundError("Repository not found")
	} else if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	return mappers.FromInventoryModelToDomain(&instance), nil
}
