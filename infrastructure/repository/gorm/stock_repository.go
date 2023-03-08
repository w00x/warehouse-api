package gorm

import (
	"gorm.io/gorm"
	"warehouse/domain"
	"warehouse/infrastructure/errors"
	"warehouse/infrastructure/repository/mappers"
	"warehouse/infrastructure/repository/models"
)

type StockRepository struct {
	postgresBase *PostgresBase
}

func NewStockRepository() *StockRepository {
	postgresBase := NewPostgresBase()
	return &StockRepository{postgresBase}
}

func (r StockRepository) All() (*[]domain.Stock, errors.IBaseError) {
	var instances []models.Stock
	result := r.postgresBase.DB.Joins("Item").Joins("Rack").Model(&models.Stock{}).Scan(&instances)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return mappers.NewStockListDomainFromModel(&instances), nil
}

func (r StockRepository) Find(id string) (*domain.Stock, errors.IBaseError) {
	var instance models.Stock
	result := r.postgresBase.DB.
		Joins("Item").
		Joins("Rack").
		First(&instance, "stocks.id = ?", id)

	if err := result.Error; err == gorm.ErrRecordNotFound {
		return nil, errors.NewNotFoundError("Stock not found")
	} else if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	return mappers.FromStockModelToDomain(&instance), nil
}

func (r StockRepository) Create(instance *domain.Stock) (*domain.Stock, errors.IBaseError) {
	model := mappers.FromStockDomainToModel(instance)
	result := r.postgresBase.DB.Create(model)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Stock not created")
	}

	return r.Find(model.ID)
}

func (r StockRepository) Update(instance *domain.Stock) (*domain.Stock, errors.IBaseError) {
	instanceModel := mappers.FromStockDomainToModel(instance)
	d, err := r.Find(instanceModel.ID)
	if err != nil {
		return nil, err
	}

	model := mappers.FromStockDomainToModel(d)
	result := r.postgresBase.DB.Model(model).Updates(instanceModel)

	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Stock not updated")
	}

	return r.Find(instance.Id())
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

func (r StockRepository) AllByInventory(inventoryId string) (*[]domain.Stock, errors.IBaseError) {
	inventoryRepo := NewInventoryRepository()
	inventory, errInventory := inventoryRepo.Find(inventoryId)

	if errInventory != nil {
		return nil, errInventory
	}

	fromDate := inventory.OperationDate
	var stocks []models.Stock

	result := r.postgresBase.DB.Joins("Item").
		Joins("Rack").
		Where("operation_date >= ?", fromDate).Find(&stocks)

	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	return mappers.NewStockListDomainFromModel(&stocks), nil
}

func (r StockRepository) AllByLastInventory() (*[]domain.Stock, errors.IBaseError) {
	inventoryRepo := NewInventoryRepository()
	inventory, errInventory := inventoryRepo.Last()

	if errInventory != nil {
		return nil, errInventory
	}

	return r.AllByInventory(inventory.Id())
}
