package repositories

import (
	"api-crud-gin/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Country interface {
	Save(models.Country) error
	Update(models.Country) error
	Delete(models.Country) error
	FindAll() ([]models.Country, error)
	FindById(string) (models.Country, error)
	CloseConnection()
}

type repo struct {
	connection *gorm.DB
}

// INSTANTIATE NEW REPOSITORY
func New() Country {
	db, err := gorm.Open("mysql", BuildURL(BuildConfig()))

	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&models.Country{}, &models.Detail{})

	return &repo{
		connection: db,
	}
}

func (r *repo) CloseConnection() {
	err := r.connection.Close()
	if err != nil {
		panic("Failed to close database")
	}
}

func (r *repo) FindAll() ([]models.Country, error) {
	var countries []models.Country
	err := r.connection.Set("gorm:auto_preload", true).Find(&countries).Error
	return countries, err
}

func (r *repo) FindById(id string) (models.Country, error) {
	var country models.Country
	err := r.connection.Set("gorm:auto_preload", true).Where("id = ?", id).First(&country).Error
	return country, err
}

func (r *repo) Save(country models.Country) error {
	return r.connection.Create(&country).Error
}

func (r *repo) Update(country models.Country) error {
	return r.connection.Save(&country).Error
}

func (r *repo) Delete(country models.Country) error {
	return r.connection.Delete(&country).Error
}
