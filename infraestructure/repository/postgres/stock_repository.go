package postgres

import (
	"gorm.io/gorm"
	"warehouse/domain"
	"warehouse/infraestructure/errors"
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
	result := r.postgresBase.DB.Model(&domain.Stock{}).Scan(&instances)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &instances, nil
}

func (r StockRepository) Find(id uint) (*domain.Stock, errors.IBaseError) {
	var instance domain.Stock
	result := r.postgresBase.DB.First(&instance, id)

	if err := result.Error; err == gorm.ErrRecordNotFound {
		return nil, errors.NewNotFoundError("Stock not found")
	} else if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	return &instance, nil
}

func (r StockRepository) Create(instance *domain.Stock) (*domain.Stock, errors.IBaseError) {
	result:= r.postgresBase.DB.Create(instance)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Stock not created")
	}

	return instance, nil
}

func (r StockRepository) Update(instance *domain.Stock) (*domain.Stock, errors.IBaseError) {
	result := r.postgresBase.DB.Save(instance)

	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Stock not updated")
	}

	return instance, nil
}

func (r StockRepository) Delete(instance *domain.Stock) errors.IBaseError {
	result := r.postgresBase.DB.Delete(instance)

	if err := result.Error; err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return errors.NewNotFoundError("Stock not deleted")
	}

	return nil
}