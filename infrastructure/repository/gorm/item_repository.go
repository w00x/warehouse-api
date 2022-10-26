package gorm

import (
	"gorm.io/gorm"
	"warehouse/domain"
	"warehouse/infrastructure/errors"
	"warehouse/infrastructure/repository/mappers"
	"warehouse/infrastructure/repository/models"
)

type ItemRepository struct {
	postgresBase *PostgresBase
}

func NewItemRepository() *ItemRepository {
	postgresBase := NewPostgresBase()
	return &ItemRepository{postgresBase}
}

func (r ItemRepository) All() (*[]domain.Item, errors.IBaseError) {
	var instances []models.Item
	result := r.postgresBase.DB.Model(&models.Item{}).Scan(&instances)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return mappers.NewItemListDomainFromModel(&instances), nil
}

func (r ItemRepository) Find(id string) (*domain.Item, errors.IBaseError) {
	var instance models.Item
	result := r.postgresBase.DB.First(&instance, "id = ?", id)

	if err := result.Error; err == gorm.ErrRecordNotFound {
		return nil, errors.NewNotFoundError("Item not found")
	} else if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	return mappers.FromItemModelToDomain(&instance), nil
}

func (r ItemRepository) Create(instance *domain.Item) (*domain.Item, errors.IBaseError) {
	model := mappers.FromItemDomainToModel(instance)
	result := r.postgresBase.DB.Create(model)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Item not created")
	}

	return r.Find(model.ID)
}

func (r ItemRepository) Update(instance *domain.Item) (*domain.Item, errors.IBaseError) {
	instanceModel := mappers.FromItemDomainToModel(instance)
	d, err := r.Find(instanceModel.ID)
	if err != nil {
		return nil, err
	}

	model := mappers.FromItemDomainToModel(d)
	result := r.postgresBase.DB.Model(model).Updates(instanceModel)

	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Item not updated")
	}

	return r.Find(instance.Id())
}

func (r ItemRepository) Delete(instance *domain.Item) errors.IBaseError {
	result := r.postgresBase.DB.Delete(mappers.FromItemDomainToModel(instance))

	if err := result.Error; err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return errors.NewNotFoundError("Item not deleted")
	}

	return nil
}
