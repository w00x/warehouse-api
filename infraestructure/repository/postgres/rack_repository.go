package postgres

import (
	"gorm.io/gorm"
	"warehouse/domain"
	"warehouse/infraestructure/errors"
)

type RackRepository struct {
	postgresBase *PostgresBase
}

func NewRackRepository() *RackRepository {
	postgresBase := NewPostgresBase()
	return &RackRepository{postgresBase}
}

func (r RackRepository) All() (*[]domain.Rack, errors.IBaseError) {
	var instances []domain.Rack
	result := r.postgresBase.DB.Model(&domain.Rack{}).Scan(&instances)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &instances, nil
}

func (r RackRepository) Find(id uint) (*domain.Rack, errors.IBaseError) {
	var instance domain.Rack
	result := r.postgresBase.DB.First(&instance, id)

	if err := result.Error; err == gorm.ErrRecordNotFound {
		return nil, errors.NewNotFoundError("Rack not found")
	} else if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	return &instance, nil
}

func (r RackRepository) Create(instance *domain.Rack) (*domain.Rack, errors.IBaseError) {
	result:= r.postgresBase.DB.Create(instance)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Rack not created")
	}

	return instance, nil
}

func (r RackRepository) Update(instance *domain.Rack) (*domain.Rack, errors.IBaseError) {
	result := r.postgresBase.DB.Save(instance)

	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Rack not updated")
	}

	return instance, nil
}

func (r RackRepository) Delete(instance *domain.Rack) errors.IBaseError {
	result := r.postgresBase.DB.Delete(instance)

	if err := result.Error; err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return errors.NewNotFoundError("Rack not deleted")
	}

	return nil
}