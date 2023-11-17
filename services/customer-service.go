package services

import (
	"BillingGo/models"
	"BillingGo/repository"
)

type CustomerService struct{}

var repo repository.BillRespository

func NewCustomerService(repoi repository.BillRespository) BillService {
	repo = repoi
	return &CustomerService{}
}

// Create implements BillService.
func (*CustomerService) Create(model *models.Customer) (*models.Customer, error) {
	//panic("unimplemented")
	return repo.CreateCutomer(model)
}

// GetAll implements BillService.
func (*CustomerService) GetAll() ([]*models.Customer, error) {
	//panic("unimplemented")
	return repo.GetAllCutomer()
}

// GetById implements BillService.
func (*CustomerService) GetById(id string) (models.Customer, error) {
	//panic("unimplemented")
	return repo.GetCutomerById(id)
}

// Update implements BillService.
func (*CustomerService) Update(model *models.Customer) (*models.Customer, error) {
	//panic("unimplemented")

	return repo.UpdateCutomer(model)
}

// Delete implements BillService.
func (*CustomerService) Delete(id string) (*models.Customer, error) {
	//panic("unimplemented")
	return repo.DeleteCutomer(id)
}
