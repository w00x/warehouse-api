package gorm

import (
	"gorm.io/gorm"
	"warehouse/domain"
	"warehouse/infrastructure/errors"
	"warehouse/infrastructure/repository/mappers"
	"warehouse/infrastructure/repository/models"
)

type RackRepository struct {
	postgresBase *PostgresBase
}

func NewRackRepository() *RackRepository {
	postgresBase := NewPostgresBase()
	return &RackRepository{postgresBase}
}

func (r RackRepository) All() (*[]domain.Rack, errors.IBaseError) {
	var instances []models.Rack
	result := r.postgresBase.DB.Model(&models.Rack{}).Scan(&instances)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return mappers.NewRackListDomainFromModel(&instances), nil
}

func (r RackRepository) Find(id string) (*domain.Rack, errors.IBaseError) {
	var instance models.Rack
	result := r.postgresBase.DB.First(&instance, "id = ?", id)

	if err := result.Error; err == gorm.ErrRecordNotFound {
		return nil, errors.NewNotFoundError("Rack not found")
	} else if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	return mappers.FromRackModelToDomain(&instance), nil
}

func (r RackRepository) Create(instance *domain.Rack) (*domain.Rack, errors.IBaseError) {
	model := mappers.FromRackDomainToModel(instance)
	result := r.postgresBase.DB.Create(model)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Rack not created")
	}

	return r.Find(model.ID)
}

func (r RackRepository) Update(instance *domain.Rack) (*domain.Rack, errors.IBaseError) {
	instanceModel := mappers.FromRackDomainToModel(instance)
	d, err := r.Find(instanceModel.ID)
	if err != nil {
		return nil, err
	}

	model := mappers.FromRackDomainToModel(d)
	result := r.postgresBase.DB.Model(model).Updates(instanceModel)

	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Rack not updated")
	}

	return r.Find(instance.Id())
}

func (r RackRepository) Delete(instance *domain.Rack) errors.IBaseError {
	result := r.postgresBase.DB.Delete(mappers.FromRackDomainToModel(instance))

	if err := result.Error; err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return errors.NewNotFoundError("Rack not deleted")
	}

	return nil
}
