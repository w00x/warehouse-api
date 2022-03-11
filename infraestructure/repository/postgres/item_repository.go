package postgres

import (
	"gorm.io/gorm"
	"warehouse/domain"
	"warehouse/infraestructure/errors"
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

	return instance.ToDomain(), nil
}

func (r ItemRepository) Create(instance *domain.Item) (*domain.Item, errors.IBaseError) {
	model := models.FromItemDomainToModel(instance)
	result:= r.postgresBase.DB.Create(model)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Item not created")
	}

	return model.ToDomain(), nil
}

func (r ItemRepository) Update(instance *domain.Item) (*domain.Item, errors.IBaseError) {
	model := models.FromItemDomainToModel(instance)
	result := r.postgresBase.DB.Save(model)

	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Item not updated")
	}

	return model.ToDomain(), nil
}

func (r ItemRepository) Delete(instance *domain.Item) errors.IBaseError {
	result := r.postgresBase.DB.Delete(models.FromItemDomainToModel(instance))

	if err := result.Error; err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return errors.NewNotFoundError("Item not deleted")
	}

	return nil
}