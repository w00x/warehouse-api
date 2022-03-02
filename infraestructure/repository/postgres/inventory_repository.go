package postgres

import (
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
	q := "SELECT id, operation_date FROM main_schema.inventories;"

	rows, err := i.postgresBase.DB.Query(q)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()
	var inventories []*domain.Inventory

	for rows.Next() {
		var inventory domain.Inventory
		rows.Scan(&inventory.Id, &inventory.OperationDate)
		inventories = append(inventories, &inventory)
	}

	return inventories, nil
}

func (i InventoryRepository) Find(id string) (*domain.Inventory, errors.IBaseError) {
	q := "SELECT id, operation_date FROM main_schema.inventories WHERE id = $1;"

	rows, err := i.postgresBase.DB.Query(q, id)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()
	var inventory domain.Inventory
	if rows.Next() {
		rows.Scan(&inventory.Id, &inventory.OperationDate)
	} else {
		return nil, errors.NewNotFoundError("Repository not found (ID="+ id +")")
	}
	return &inventory, nil
}

func (i InventoryRepository) Update(id string, operationDate time.Time) errors.IBaseError {
	q := `UPDATE main_schema.inventories
	SET operation_date = $2
	WHERE id = $1;`

	rows, err := i.postgresBase.DB.Exec(q, id, operationDate)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	count, err := rows.RowsAffected()
	if err != nil {
		panic(err)
	}

	if count == 0 {
		return errors.NewNotFoundError("Repository not found (ID="+ id +")")
	}

	return nil
}

func (i InventoryRepository) Create(operationDate time.Time) (*domain.Inventory, errors.IBaseError) {
	q := `INSERT INTO main_schema.inventories (operation_date)
		  VALUES ($1) RETURNING id;`

	rows, err := i.postgresBase.DB.Query(q, operationDate)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	var inventory domain.Inventory
	if rows.Next() {
		rows.Scan(&inventory.Id)
		inventory.OperationDate = operationDate
	} else {
		return nil, errors.NewNotCreatedError("Repository not created")
	}
	return &inventory, nil
}

func (i InventoryRepository) Delete(id string) errors.IBaseError {
	q := `DELETE FROM main_schema.inventories WHERE id = $1;`

	rows, err := i.postgresBase.DB.Exec(q, id)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	count, err := rows.RowsAffected()
	if err != nil {
		panic(err)
	}

	if count == 0 {
		return errors.NewNotFoundError("Repository not found (ID="+ id +")")
	}

	return nil
}
