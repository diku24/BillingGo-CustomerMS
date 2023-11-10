package repository

import "BillingGo/models"

type BillRespository interface {
	CreateCutomer(model *models.Customer) (*models.Customer, error)
	GetAllCutomer() ([]models.Customer, error)
	GetCutomerById(id string) (models.Customer, error)
	UpdateCutomer(id string) error
	DeleteCutomer(id string) error
	PingServer()
}
