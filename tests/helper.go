package tests

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	localGorm "warehouse/infrastructure/repository/gorm"
	"warehouse/shared"
)

func SetDbConn() {
	uri := shared.GetEnv("DATABASE_TEST_URI")
	var err error
	localGorm.DbConn, err = gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func GetCurrentAdapter() string {
	return "gorm"
}

func CleanInventory() {
	localGorm.NewPostgresBase().DB.Exec("DELETE FROM inventories")
}

func CleanItem() {
	localGorm.NewPostgresBase().DB.Exec("DELETE FROM items")
}

func CleanMarket() {
	localGorm.NewPostgresBase().DB.Exec("DELETE FROM markets")
}

func CleanPrice() {
	localGorm.NewPostgresBase().DB.Exec("DELETE FROM prices")
}

func CleanRack() {
	localGorm.NewPostgresBase().DB.Exec("DELETE FROM racks")
}

func CleanStock() {
	localGorm.NewPostgresBase().DB.Exec("DELETE FROM stocks")
}
