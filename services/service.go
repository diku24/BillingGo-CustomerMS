package services

import (
	"CustomerMS/models"
)

type BillService interface {
	Create(model *models.Customer) (*models.Customer, error)
	GetAll() ([]*models.Customer, error)
	GetById(id string) (models.Customer, error)
	Update(model *models.Customer) (*models.Customer, error)
	//Update(id string) (*models.Customer, error)
	Delete(id string) (*models.Customer, error)
}
