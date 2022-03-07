package postgres

import (
	"gorm.io/gorm"
	"time"
	"warehouse/domain"
	"warehouse/infraestructure/errors"
)

type InventoryRepository struct {
	postgresBase *PostgresBase
}

func NewInventoryRepository() *InventoryRepository {
	postgresBase := NewPostgresBase()
	return &InventoryRepository{postgresBase}
}

func (i InventoryRepository) All() ([]*domain.Inventory, errors.IBaseError) {
	rows, err := i.postgresBase.DB.Model(&domain.Inventory{}).Rows()
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()
	var inventories []*domain.Inventory

	for rows.Next() {
		var inventory domain.Inventory
		err := rows.Scan(&inventory.Id, &inventory.OperationDate, &inventory.CreatedAt, &inventory.UpdatedAt)
		if err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		inventories = append(inventories, &inventory)
	}

	return inventories, nil
}

func (i InventoryRepository) Find(id string) (*domain.Inventory, errors.IBaseError) {
	var inventory domain.Inventory
	result := i.postgresBase.DB.First(&inventory, id)

	if err := result.Error; err == gorm.ErrRecordNotFound {
		return nil, errors.NewNotFoundError("Repository not found")
	} else if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	return &inventory, nil
}

func (i InventoryRepository) Create(operationDate time.Time) (*domain.Inventory, errors.IBaseError) {
	inventory := domain.NewInventory(0, operationDate)
	result:= i.postgresBase.DB.Create(inventory)
	if err := result.Error; err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return nil, errors.NewNotFoundError("Repository not created")
	}

	return inventory, nil
}

func (i InventoryRepository) Update(id string, operationDate time.Time) errors.IBaseError {
	var inventory domain.Inventory
	result := i.postgresBase.DB.First(&inventory, id)

	if err := result.Error; err == gorm.ErrRecordNotFound {
		return errors.NewNotFoundError("Repository not found")
	} else if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	inventory.OperationDate = operationDate
	result = i.postgresBase.DB.Save(inventory)

	if err := result.Error; err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return errors.NewNotFoundError("Repository not updated")
	}

	return nil
}

func (i InventoryRepository) Delete(id string) errors.IBaseError {
	var inventory domain.Inventory
	result := i.postgresBase.DB.First(&inventory, id)

	if err := result.Error; err == gorm.ErrRecordNotFound {
		return errors.NewNotFoundError("Repository not found")
	} else if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	result = i.postgresBase.DB.Delete(inventory)

	if err := result.Error; err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	count := result.RowsAffected

	if count == 0 {
		return errors.NewNotFoundError("Repository not deleted")
	}

	return nil
}
