package postgres

import (
	"gorm.io/gorm"
	"warehouse/domain"
	"warehouse/infraestructure/errors"
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
	result := r.postgresBase.DB.Model(&domain.Item{}).Scan(&instances)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &instances, nil
}

func (r ItemRepository) Find(id uint) (*domain.Item, errors.IBaseError) {
	var instance domain.Item
	result := r.postgresBase.DB.First(&instance, id)

	if err := result.Error; err == gorm.ErrRecordNotFound {
		return nil, errors.NewNotFoundError("Item not found")
	} else if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	return &instance, nil
}

func (r ItemRepository) Create(instance *domain.Item) (*domain.Item, errors.IBaseError) {
	result:= r.postgresBase.DB.Create(instance)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Item not created")
	}

	return instance, nil
}

func (r ItemRepository) Update(instance *domain.Item) (*domain.Item, errors.IBaseError) {
	result := r.postgresBase.DB.Save(instance)

	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Item not updated")
	}

	return instance, nil
}

func (r ItemRepository) Delete(instance *domain.Item) errors.IBaseError {
	result := r.postgresBase.DB.Delete(instance)

	if err := result.Error; err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return errors.NewNotFoundError("Item not deleted")
	}

	return nil
}