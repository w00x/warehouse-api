package postgres

import (
	"gorm.io/gorm"
	"warehouse/domain"
	"warehouse/infraestructure/errors"
)

type MarketRepository struct {
	postgresBase *PostgresBase
}

func NewMarketRepository() *MarketRepository {
	postgresBase := NewPostgresBase()
	return &MarketRepository{postgresBase}
}

func (r MarketRepository) All() (*[]domain.Market, errors.IBaseError) {
	var instances []domain.Market
	result := r.postgresBase.DB.Model(&domain.Market{}).Scan(&instances)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &instances, nil
}

func (r MarketRepository) Find(id uint) (*domain.Market, errors.IBaseError) {
	var instance domain.Market
	result := r.postgresBase.DB.First(&instance, id)

	if err := result.Error; err == gorm.ErrRecordNotFound {
		return nil, errors.NewNotFoundError("Market not found")
	} else if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	return &instance, nil
}

func (r MarketRepository) Create(instance *domain.Market) (*domain.Market, errors.IBaseError) {
	result:= r.postgresBase.DB.Create(instance)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Market not created")
	}

	return instance, nil
}

func (r MarketRepository) Update(instance *domain.Market) (*domain.Market, errors.IBaseError) {
	result := r.postgresBase.DB.Save(instance)

	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Market not updated")
	}

	return instance, nil
}

func (r MarketRepository) Delete(instance *domain.Market) errors.IBaseError {
	result := r.postgresBase.DB.Delete(instance)

	if err := result.Error; err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return errors.NewNotFoundError("Market not deleted")
	}

	return nil
}