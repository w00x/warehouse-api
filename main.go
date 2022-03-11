package main

import "warehouse/infraestructure"

func main() {
	infraestructure.RunMigrations()
	infraestructure.Routes()
}
