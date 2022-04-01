package factories

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"testing"
	"time"
	"warehouse/domain"
	"warehouse/infraestructure/repository/postgres"
	"warehouse/shared"
)

type Inventory struct {
	OperationDate 	shared.DateTime
}

func (i Inventory) ToDomain() *domain.Inventory {
	return domain.NewInventory(0, i.OperationDate)
}

func NewInventoryFactory(t *testing.T) *domain.Inventory {
	inventory := &Inventory{}
	err := gofakeit.Struct(inventory)
	if err != nil {
		fmt.Println(err)
	}

	repo := postgres.NewInventoryRepository()
	inventoryDomain, errRepo := repo.Create(inventory.ToDomain())
	if errRepo != nil {
		panic(err)
	}

	t.Cleanup(func() {
		CleanInventory()
	})

	return inventoryDomain
}

func NewInventoryObjectFactory() map[string]interface{} {
	inventory := &Inventory{}
	err := gofakeit.Struct(inventory)
	if err != nil {
		fmt.Println(err)
	}

	inventoryMarshal := map[string]interface{}{
		"operation_date": time.Time(inventory.OperationDate).Format("2006-01-02 15:04:05"),
	}

	return inventoryMarshal
}

func NewInventoryFactoryList(count int, t *testing.T) []*domain.Inventory {
	var inventoryDomains []*domain.Inventory
	repo := postgres.NewInventoryRepository()

	for i := 0; i < count; i++ {
		inventory := &Inventory{}
		err := gofakeit.Struct(inventory)
		if err != nil {
			panic(err)
		}

		inventoryDomain, errRepo := repo.Create(inventory.ToDomain())
		if errRepo != nil {
			panic(err)
		}
		inventoryDomains = append(inventoryDomains, inventoryDomain)
	}

	t.Cleanup(func() {
		CleanInventory()
	})

	return inventoryDomains
}

func CleanInventory() {
	postgres.NewPostgresBase().DB.Exec("DELETE FROM inventories")
	postgres.NewPostgresBase().DB.Exec("ALTER SEQUENCE inventories_id_seq RESTART WITH 1")
}
