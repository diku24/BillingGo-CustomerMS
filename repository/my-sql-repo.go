package repository

import (
	repository "BillingGo/db"
	"BillingGo/models"
	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MySQLRepository struct{}

var (
	customer models.Customer
)

func NewMySQLReopsitory() BillRespository {
	return &MySQLRepository{}
}

// CreateCutomer implements BillRespository.
func (*MySQLRepository) CreateCutomer(model *models.Customer) (*models.Customer, error) {

	db, err := repository.OpenMysqlConnection()

	customerDataHolder := models.Customer{
		CustomerId:    model.CustomerId,
		CustomerName:  model.CustomerName,
		ContactNumber: model.ContactNumber,
		Address:       model.Address,
		Priority:      model.Priority,
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
		DeletedAt:     model.DeletedAt,
	}

	db.Create(&customerDataHolder)
	// defer db.Close()
	return &customerDataHolder, err

}

// GetCutomerById implements BillRespository.
func (*MySQLRepository) GetCutomerById(id string) (models.Customer, error) {

	db, err := repository.OpenMysqlConnection()
	if err != nil {
		logrus.Errorln(err.Error())
		logrus.Errorln("Failed to Connect to the Database !!")
	}

	logrus.Println("You have entered Id: " + id)
	result := db.First(&customer, id)
	logrus.Println(result.RowsAffected)
	logrus.Errorln(result.Error)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		logrus.Errorf("No Records Found for Entered Id: %v", id)
	}

	logrus.Println("Values "+customer.CustomerId, customer.CustomerName, customer.CreatedAt)
	return customer, nil

}

// DeleteCutomer implements BillRespository.
func (*MySQLRepository) DeleteCutomer(id string) error {
	panic("unimplemented")

}

// GetAllCutomer implements BillRespository.
func (*MySQLRepository) GetAllCutomer() ([]models.Customer, error) {
	panic("unimplemented")
}

// UpdateCutomer implements BillRespository.
func (*MySQLRepository) UpdateCutomer(id string) error {
	panic("unimplemented")
}
