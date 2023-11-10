package services

import (
	"BillingGo/models"
)

type BillService interface {
	Create(models *models.Customer) (*models.Customer, error)
	GetAll() ([]models.Customer, error)
	GetById(id string) (models.Customer, error)
	Update(id string) error
	Delete(id string) error
}