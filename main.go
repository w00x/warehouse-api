package main

import (
	"os"
	"warehouse/infraestructure"
)

func main() {
	os.Setenv("TZ", "UTC-0")
	infraestructure.RunMigrations()
	infraestructure.Routes()
}
