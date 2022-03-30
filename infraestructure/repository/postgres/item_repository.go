package postgres

import (
	"gorm.io/gorm"
	"warehouse/domain"
	"warehouse/infraestructure/errors"
	"warehouse/infraestructure/repository/mappers"
	"warehouse/infraestructure/repository/models"
)

type ItemRepository struct {
	postgresBase *PostgresBase
}

func NewItemRepository() *ItemRepository {
	postgresBase := NewPostgresBase()
	return &ItemRepository{postgresBase}
}

func (r ItemRepository) All() (*[]domain.Item, errors.IBaseError) {
	var instances []domain.Item
	result := r.postgresBase.DB.Model(&models.Item{}).Scan(&instances)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &instances, nil
}

func (r ItemRepository) Find(id uint) (*domain.Item, errors.IBaseError) {
	var instance models.Item
	result := r.postgresBase.DB.First(&instance, id)

	if err := result.Error; err == gorm.ErrRecordNotFound {
		return nil, errors.NewNotFoundError("Item not found")
	} else if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	return mappers.FromItemModelToDomain(&instance), nil
}

func (r ItemRepository) Create(instance *domain.Item) (*domain.Item, errors.IBaseError) {
	model := mappers.FromItemDomainToModel(instance)
	result:= r.postgresBase.DB.Create(model)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Item not created")
	}

	return r.Find(model.Id)
}

func (r ItemRepository) Update(instance *domain.Item) (*domain.Item, errors.IBaseError) {
	d, err := r.Find(instance.Id)
	if err != nil {
		return nil, err
	}

	result := r.postgresBase.DB.Model(d).Updates(instance)

	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Item not updated")
	}

	return r.Find(instance.Id)
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