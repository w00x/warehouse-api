package infraestructure

import (
	"warehouse/infraestructure/repository/models"
	"warehouse/infraestructure/repository/postgres"
)

func RunMigrations() {
	db := postgres.NewPostgresBase()
	db.DB.AutoMigrate(models.Inventory{})
	db.DB.AutoMigrate(models.Rack{})
	db.DB.AutoMigrate(models.Market{})
	db.DB.AutoMigrate(models.Item{})
	db.DB.AutoMigrate(models.Price{})
	db.DB.AutoMigrate(models.Stock{})
}
