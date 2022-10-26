package injectors

import (
	"warehouse/application"
	domainRepository "warehouse/domain/repository"
	v1 "warehouse/infrastructure/controller/v1"
	"warehouse/infrastructure/errors"
	"warehouse/infrastructure/factory/repository"
	localGorm "warehouse/infrastructure/repository/gorm"
	"warehouse/tests"
)

func InventoryControllerInjector() *v1.InventoryController {
	return v1.NewInventoryController(InventoryApplicationInjector())
}

func InventoryApplicationInjector() *application.InventoryApplication {
	factory, err := InventoryRepository()
	if err != nil {
		panic(err)
	}
	return application.NewInventoryApplication(factory)
}

func InventoryRepository() (domainRepository.IInventoryRepository, errors.IBaseError) {
	return repository.InventoryFactory(tests.GetCurrentAdapter())
}

func CleanInventory() {
	localGorm.NewPostgresBase().DB.Exec("DELETE FROM inventories")
}
