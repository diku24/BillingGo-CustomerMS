package repository

import "CustomerMS/models"

type BillRespository interface {
	CreateCutomer(model *models.Customer) (*models.Customer, error)
	GetAllCutomer() ([]*models.Customer, error)
	GetCutomerById(id string) (models.Customer, error)
	UpdateCutomer(model *models.Customer) (*models.Customer, error)
	//UpdateCutomer(id string) (*models.Customer, error)
	DeleteCutomer(id string) (*models.Customer, error)
}
