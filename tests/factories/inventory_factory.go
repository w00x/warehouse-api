package factories

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"warehouse/domain"
	"warehouse/infrastructure/repository/gorm"
	"warehouse/shared"
	"warehouse/tests/injectors"
)

type Inventory struct {
	OperationDate shared.DateTime
}

func (i Inventory) ToDomain() *domain.Inventory {
	return domain.NewInventory("", i.OperationDate)
}

func NewInventoryFactory() *domain.Inventory {
	inventory := &Inventory{}
	err := gofakeit.Struct(inventory)
	if err != nil {
		fmt.Println(err)
	}

	repo, _ := injectors.InventoryRepository()
	inventoryDomain, errRepo := repo.Create(inventory.ToDomain())
	if errRepo != nil {
		panic(errRepo)
	}

	return inventoryDomain
}

func NewInventoryDomainFactory() *domain.Inventory {
	inventory := &Inventory{}
	err := gofakeit.Struct(inventory)
	if err != nil {
		fmt.Println(err)
	}

	return inventory.ToDomain()
}

func NewInventoryObjectFactory() map[string]interface{} {
	inventory := &Inventory{}
	err := gofakeit.Struct(inventory)
	if err != nil {
		fmt.Println(err)
	}

	inventoryMarshal := map[string]interface{}{
		"operation_date": inventory.OperationDate.Format("2006-01-02 15:04:05"),
	}

	return inventoryMarshal
}

func NewInventoryFactoryList(count int) []*domain.Inventory {
	var inventoryDomains []*domain.Inventory
	repo := gorm.NewInventoryRepository()

	for i := 0; i < count; i++ {
		inventory := &Inventory{}
		err := gofakeit.Struct(inventory)
		if err != nil {
			panic(err)
		}

		inventoryDomain, errRepo := repo.Create(inventory.ToDomain())
		if errRepo != nil {
			panic(errRepo)
		}
		inventoryDomains = append(inventoryDomains, inventoryDomain)
	}

	return inventoryDomains
}
