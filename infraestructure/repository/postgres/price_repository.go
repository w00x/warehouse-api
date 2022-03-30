package postgres

import (
	"gorm.io/gorm"
	"warehouse/domain"
	"warehouse/infraestructure/errors"
	"warehouse/infraestructure/repository/mappers"
	"warehouse/infraestructure/repository/models"
)

type PriceRepository struct {
	postgresBase *PostgresBase
}

func NewPriceRepository() *PriceRepository {
	postgresBase := NewPostgresBase()
	return &PriceRepository{postgresBase}
}

func (r PriceRepository) All() (*[]domain.Price, errors.IBaseError) {
	var instances []domain.Price
	result := r.postgresBase.DB.Joins("Market").Joins("Item").Model(&models.Price{}).Scan(&instances)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &instances, nil
}

func (r PriceRepository) Find(id uint) (*domain.Price, errors.IBaseError) {
	var instance models.Price
	result := r.postgresBase.DB.Joins("Market").Joins("Item").First(&instance, id)

	if err := result.Error; err == gorm.ErrRecordNotFound {
		return nil, errors.NewNotFoundError("Price not found")
	} else if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	return mappers.FromPriceModelToDomain(&instance), nil
}

func (r PriceRepository) Create(instance *domain.Price) (*domain.Price, errors.IBaseError) {
	model := mappers.FromPriceDomainToModel(instance)
	result:= r.postgresBase.DB.Create(model)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Price not created")
	}

	return r.Find(model.Id)
}

func (r PriceRepository) Update(instance *domain.Price) (*domain.Price, errors.IBaseError) {
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
		return nil, errors.NewNotFoundError("Price not updated")
	}

	return r.Find(instance.Id)
}

func (r PriceRepository) Delete(instance *domain.Price) errors.IBaseError {
	result := r.postgresBase.DB.Delete(mappers.FromPriceDomainToModel(instance))

	if err := result.Error; err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return errors.NewNotFoundError("Price not deleted")
	}

	return nil
}