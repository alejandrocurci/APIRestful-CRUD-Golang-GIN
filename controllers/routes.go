package controllers

import (
	"first-api/services"

	"github.com/gin-gonic/gin"
)

// Configure routes
func SetupRoutes() *gin.Engine {
	r := gin.Default()
	group := r.Group("/countries-api")
	{
		group.GET("/countries", services.GetCountries)
		group.GET("/country/:id", services.GetCountryByID)
		group.POST("/country", services.CreateCountry)
		group.PUT("/country/:id", services.UpdateCountry)
		group.DELETE("/country/:id", services.DeleteCountry)
	}
	return r
}
