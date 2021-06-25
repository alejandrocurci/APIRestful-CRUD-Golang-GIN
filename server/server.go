package server

import (
	"api-crud-gin/controllers"
	"api-crud-gin/repositories"
	"api-crud-gin/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	repository repositories.Country
	service    services.Country
	controller controllers.Country
)

// Build function initializes the server
func Build() *gin.Engine {

	r := gin.New()

	repository = repositories.New()
	service = services.New(repository)
	controller = controllers.New(service)

	group := r.Group("/api")
	{
		group.GET("/countries", GetAll)
		group.GET("/country/:id", GetById)
		group.POST("/country", Create)
		group.PUT("/country/:id", Update)
		group.DELETE("/country/:id", Delete)
	}

	return r
}

// GetAll handler
func GetAll(ctx *gin.Context) {
	countries, err := controller.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, countries)
}

// GetById handler
func GetById(ctx *gin.Context) {
	country, err := controller.FindById(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, country)
}

// Create handler
func Create(ctx *gin.Context) {

	if err := controller.Save(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success!"})
}

// Update handler
func Update(ctx *gin.Context) {

	if err := controller.Update(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success!"})
}

// Delete handler
func Delete(ctx *gin.Context) {

	if err := controller.Delete(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success!"})
}
