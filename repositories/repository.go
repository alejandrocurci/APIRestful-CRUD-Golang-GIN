package repositories

import (
	"first-api/config"
	"first-api/models"

	_ "github.com/go-sql-driver/mysql"
)

// GET ALL COUNTRIES
func GetAllCountries(countries *[]models.Country) (err error) {
	if err := config.DB.Find(countries).Error; err != nil {
		return err
	}
	return nil
}

// GET COUNTRY BY ID
func GetCountryById(country *models.Country, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(country).Error; err != nil {
		return err
	}
	return nil
}

// CREATE COUNTRY
func CreateCountry(country *models.Country) (err error) {
	if err := config.DB.Create(country).Error; err != nil {
		return err
	}
	return nil
}

// UPDATE COUNTRY
func UpdateCountry(country *models.Country, id string) (err error) {
	config.DB.Save(country)
	return nil
}

// DELETE COUNTRY
func DeleteCountry(country *models.Country, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(country).Error; err != nil {
		return err
	}
	config.DB.Where("id = ?", id).Delete(country)
	return nil
}
