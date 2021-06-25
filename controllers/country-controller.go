package controllers

import (
	"api-crud-gin/models"
	"api-crud-gin/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Country interface {
	Save(*gin.Context) error
	Update(*gin.Context) error
	Delete(*gin.Context) error
	FindAll() ([]models.Country, error)
	FindById(*gin.Context) (models.Country, error)
}

type controller struct {
	service services.Country
}

var validate *validator.Validate

func New(s services.Country) Country {
	validate = validator.New()
	return &controller{
		service: s,
	}
}

func (c *controller) FindAll() ([]models.Country, error) {
	return c.service.FindAll()
}

func (c *controller) FindById(ctx *gin.Context) (models.Country, error) {
	id := ctx.Params.ByName("id")
	return c.service.FindById(id)
}

func (c *controller) Save(ctx *gin.Context) error {
	var country models.Country
	err := ctx.ShouldBindJSON(&country)
	if err != nil {
		return err
	}
	err = validate.Struct(country)
	if err != nil {
		return err
	}
	c.service.Save(country)
	return nil
}

func (c *controller) Update(ctx *gin.Context) error {
	var country models.Country
	err := ctx.ShouldBindJSON(&country)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	country.ID = uint(id)

	err = validate.Struct(country)
	if err != nil {
		return err
	}
	c.service.Update(country)
	return nil
}

func (c *controller) Delete(ctx *gin.Context) error {
	var country models.Country
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	country.ID = uint(id)
	c.service.Delete(country)
	return nil
}
