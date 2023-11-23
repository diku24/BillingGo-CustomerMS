package services

import (
	"BillingGo/models"
	"BillingGo/repository"
)

type CustomerService struct {
	repository repository.BillRespository
}

func NewCustomerService(repoi repository.BillRespository) BillService {
	return &CustomerService{
		repository: repoi,
	}
}

// Create implements BillService.
func (r *CustomerService) Create(model *models.Customer) (*models.Customer, error) {
	//panic("unimplemented")
	//return repo.CreateCutomer(model)
	return r.repository.CreateCutomer(model)
}

// GetAll implements BillService.
func (r *CustomerService) GetAll() ([]*models.Customer, error) {
	//panic("unimplemented")
	return r.repository.GetAllCutomer()
}

// GetById implements BillService.
func (r *CustomerService) GetById(id string) (models.Customer, error) {
	//panic("unimplemented")
	return r.repository.GetCutomerById(id)
}

// Update implements BillService.
func (r *CustomerService) Update(model *models.Customer) (*models.Customer, error) {
	//panic("unimplemented")

	return r.repository.UpdateCutomer(model)
}

// Delete implements BillService.
func (r *CustomerService) Delete(id string) (*models.Customer, error) {
	//panic("unimplemented")
	return r.repository.DeleteCutomer(id)
}
