package main

import (
	"first-api/config"
	"first-api/controllers"
	"first-api/models"
	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

func main() {

	config.DB, err = gorm.Open("mysql", config.BuildURL(config.BuildConfig()))

	if err != nil {
		fmt.Println("Status: ", err)
	}

	defer config.DB.Close()

	config.DB.AutoMigrate(&models.Country{})

	r := controllers.SetupRoutes()
	r.Run(":8080")
}
