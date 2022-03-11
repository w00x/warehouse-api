package postgres

import (
	"gorm.io/gorm"
	"warehouse/domain"
	"warehouse/infraestructure/errors"
	"warehouse/infraestructure/repository/models"
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
	result := r.postgresBase.DB.Model(&models.Market{}).Scan(&instances)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &instances, nil
}

func (r MarketRepository) Find(id uint) (*domain.Market, errors.IBaseError) {
	var instance models.Market
	result := r.postgresBase.DB.First(&instance, id)

	if err := result.Error; err == gorm.ErrRecordNotFound {
		return nil, errors.NewNotFoundError("Market not found")
	} else if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	return instance.ToDomain(), nil
}

func (r MarketRepository) Create(instance *domain.Market) (*domain.Market, errors.IBaseError) {
	model := models.FromMarketDomainToModel(instance)
	result:= r.postgresBase.DB.Create(model)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Market not created")
	}

	return model.ToDomain(), nil
}

func (r MarketRepository) Update(instance *domain.Market) (*domain.Market, errors.IBaseError) {
	model := models.FromMarketDomainToModel(instance)
	result := r.postgresBase.DB.Save(model)

	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Market not updated")
	}

	return model.ToDomain(), nil
}

func (r MarketRepository) Delete(instance *domain.Market) errors.IBaseError {
	result := r.postgresBase.DB.Delete(models.FromMarketDomainToModel(instance))

	if err := result.Error; err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return errors.NewNotFoundError("Market not deleted")
	}

	return nil
}