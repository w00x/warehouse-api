package postgres

import (
	"gorm.io/gorm"
	"warehouse/domain"
	"warehouse/infraestructure/errors"
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
	result := r.postgresBase.DB.Model(&domain.Price{}).Scan(&instances)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &instances, nil
}

func (r PriceRepository) Find(id uint) (*domain.Price, errors.IBaseError) {
	var instance domain.Price
	result := r.postgresBase.DB.First(&instance, id)

	if err := result.Error; err == gorm.ErrRecordNotFound {
		return nil, errors.NewNotFoundError("Price not found")
	} else if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	return &instance, nil
}

func (r PriceRepository) Create(instance *domain.Price) (*domain.Price, errors.IBaseError) {
	result:= r.postgresBase.DB.Create(instance)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Price not created")
	}

	return instance, nil
}

func (r PriceRepository) Update(instance *domain.Price) (*domain.Price, errors.IBaseError) {
	result := r.postgresBase.DB.Save(instance)

	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Price not updated")
	}

	return instance, nil
}

func (r PriceRepository) Delete(instance *domain.Price) errors.IBaseError {
	result := r.postgresBase.DB.Delete(instance)

	if err := result.Error; err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return errors.NewNotFoundError("Price not deleted")
	}

	return nil
}