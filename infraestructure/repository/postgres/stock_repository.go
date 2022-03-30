package postgres

import (
	"gorm.io/gorm"
	"warehouse/domain"
	"warehouse/infraestructure/errors"
	"warehouse/infraestructure/repository/mappers"
	"warehouse/infraestructure/repository/models"
)

type StockRepository struct {
	postgresBase *PostgresBase
}

func NewStockRepository() *StockRepository {
	postgresBase := NewPostgresBase()
	return &StockRepository{postgresBase}
}

func (r StockRepository) All() (*[]domain.Stock, errors.IBaseError) {
	var instances []domain.Stock
	result := r.postgresBase.DB.Joins("Item").Joins("Rack").Model(&models.Stock{}).Scan(&instances)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &instances, nil
}

func (r StockRepository) Find(id uint) (*domain.Stock, errors.IBaseError) {
	var instance models.Stock
	result := r.postgresBase.DB.Joins("Item").Joins("Rack").First(&instance, id)

	if err := result.Error; err == gorm.ErrRecordNotFound {
		return nil, errors.NewNotFoundError("Stock not found")
	} else if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	return mappers.FromStockModelToDomain(&instance), nil
}

func (r StockRepository) Create(instance *domain.Stock) (*domain.Stock, errors.IBaseError) {
	model := mappers.FromStockDomainToModel(instance)
	result:= r.postgresBase.DB.Create(model)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Stock not created")
	}

	return r.Find(model.Id)
}

func (r StockRepository) Update(instance *domain.Stock) (*domain.Stock, errors.IBaseError) {
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
		return nil, errors.NewNotFoundError("Stock not updated")
	}

	return r.Find(instance.Id)
}

func (r StockRepository) Delete(instance *domain.Stock) errors.IBaseError {
	result := r.postgresBase.DB.Delete(mappers.FromStockDomainToModel(instance))

	if err := result.Error; err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return errors.NewNotFoundError("Stock not deleted")
	}

	return nil
}