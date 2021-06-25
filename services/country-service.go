package services

import (
	"api-crud-gin/models"
	"api-crud-gin/repositories"
)

type Country interface {
	Save(models.Country) error
	Update(models.Country) error
	Delete(models.Country) error
	FindAll() ([]models.Country, error)
	FindById(string) (models.Country, error)
}

type service struct {
	repository repositories.Country
}

func New(repo repositories.Country) Country {
	return &service{
		repository: repo,
	}
}

func (s *service) FindAll() ([]models.Country, error) {
	return s.repository.FindAll()
}

func (s *service) FindById(id string) (models.Country, error) {
	return s.repository.FindById(id)
}

func (s *service) Save(country models.Country) error {
	return s.repository.Save(country)
}

func (s *service) Update(country models.Country) error {
	return s.repository.Update(country)
}

func (s *service) Delete(country models.Country) error {
	return s.repository.Delete(country)
}
