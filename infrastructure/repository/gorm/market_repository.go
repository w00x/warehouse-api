package gorm

import (
	"gorm.io/gorm"
	"warehouse/domain"
	"warehouse/infrastructure/errors"
	"warehouse/infrastructure/repository/mappers"
	"warehouse/infrastructure/repository/models"
)

type MarketRepository struct {
	postgresBase *PostgresBase
}

func NewMarketRepository() *MarketRepository {
	postgresBase := NewPostgresBase()
	return &MarketRepository{postgresBase}
}

func (r MarketRepository) All() (*[]domain.Market, errors.IBaseError) {
	var instances []models.Market
	result := r.postgresBase.DB.Model(&models.Market{}).Scan(&instances)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return mappers.NewMarketListDomainFromModel(&instances), nil
}

func (r MarketRepository) Find(id string) (*domain.Market, errors.IBaseError) {
	var instance models.Market
	result := r.postgresBase.DB.First(&instance, "id = ?", id)

	if err := result.Error; err == gorm.ErrRecordNotFound {
		return nil, errors.NewNotFoundError("Market not found")
	} else if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	return mappers.FromMarketModelToDomain(&instance), nil
}

func (r MarketRepository) Create(instance *domain.Market) (*domain.Market, errors.IBaseError) {
	model := mappers.FromMarketDomainToModel(instance)
	result := r.postgresBase.DB.Create(model)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Market not created")
	}

	return r.Find(model.ID)
}

func (r MarketRepository) Update(instance *domain.Market) (*domain.Market, errors.IBaseError) {
	instanceModel := mappers.FromMarketDomainToModel(instance)
	d, err := r.Find(instanceModel.ID)
	if err != nil {
		return nil, err
	}

	model := mappers.FromMarketDomainToModel(d)
	result := r.postgresBase.DB.Model(model).Updates(instanceModel)

	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Market not updated")
	}

	return r.Find(instance.Id())
}

func (r MarketRepository) Delete(instance *domain.Market) errors.IBaseError {
	result := r.postgresBase.DB.Delete(mappers.FromMarketDomainToModel(instance))

	if err := result.Error; err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return errors.NewNotFoundError("Market not deleted")
	}

	return nil
}
