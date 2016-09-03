package main

import (
	"fmt"
	"log"

	"github.com/alexyslozada/accounting-go/controller"
	"github.com/alexyslozada/accounting-go/models"
)

func main() {
	profile := models.Profile{Name: "Desde Go 3"}
	err := controller.ProfileInsert(&profile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(profile)
}
