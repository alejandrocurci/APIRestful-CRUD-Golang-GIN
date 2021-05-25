package services

import (
	"first-api/models"
	"first-api/repositories"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET ALL COUNTRIES
func GetCountries(c *gin.Context) {
	var countries []models.Country
	err := repositories.GetAllCountries(&countries)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, countries)
	}
}

// GET COUNTRY BY ID
func GetCountryByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var country models.Country
	err := repositories.GetCountryById(&country, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
	} else {
		c.JSON(http.StatusOK, country)
	}
}

// CREATE COUNTRY
func CreateCountry(c *gin.Context) {
	var country models.Country
	c.BindJSON(&country)
	err := repositories.CreateCountry(&country)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, country)
	}
}

// UPDATE COUNTRY
func UpdateCountry(c *gin.Context) {
	var country models.Country
	id := c.Params.ByName("id")
	err := repositories.GetCountryById(&country, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
	} else {
		c.BindJSON(&country)
		err = repositories.UpdateCountry(&country, id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not updated!"})
		} else {
			c.JSON(http.StatusOK, country)
		}
	}
}

// DELETE COUNTRY
func DeleteCountry(c *gin.Context) {
	var country models.Country
	id := c.Params.ByName("id")
	err := repositories.DeleteCountry(&country, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
	} else {
		c.JSON(http.StatusOK, gin.H{"id = " + id: "Successfully deleted"})
	}
}
