package main

import (
	"warehouse/infrastructure"
)

func main() {
	infrastructure.Routes("gorm").Run()
}
