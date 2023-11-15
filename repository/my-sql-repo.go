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
	result := db.First(&customer, "customer_id= ?", id)
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

	db, err := repository.OpenMysqlConnection()
	if err != nil {
		logrus.Errorln(err.Error())
		logrus.Errorln("Failed to Connect to the Database !!")
		return err
	}

	logrus.Println("You have entered Id: " + id)
	result := db.Delete(&customer, "customer_id= ?", id)
	logrus.Println(result.RowsAffected)
	logrus.Errorln(result.Error)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		logrus.Errorf("No Records Found for Entered Id: %v", id)
	}

	logrus.Println("Deleted Values "+customer.CustomerId, customer.CustomerName, customer.CreatedAt, customer.DeletedAt)

	return nil

}

// GetAllCutomer implements BillRespository.
func (*MySQLRepository) GetAllCutomer() ([]*models.Customer, error) {
	var customers []*models.Customer

	db, err := repository.OpenMysqlConnection()
	if err != nil {
		return nil, err
	}
	// Get all records
	result := db.Find(&customers) // SELECT * FROM customers;

	logrus.Printf("No of Rows affected: %d", result.RowsAffected) // returns found records count, equals `len(users)`
	logrus.Error(result.Error)                                    // returns error

	// check error ErrRecordNotFound
	errors.Is(result.Error, gorm.ErrRecordNotFound)

	return customers, nil
}

// UpdateCutomer implements BillRespository.
func (*MySQLRepository) UpdateCutomer(model *models.Customer) (*models.Customer, error) {
	db, err := repository.OpenMysqlConnection()
	if err != nil {
		return nil, err
	}

	customerDataHolder := models.Customer{
		CustomerName:  model.CustomerName,
		ContactNumber: model.ContactNumber,
		Address:       model.Address,
		Priority:      model.Priority,
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
		DeletedAt:     model.DeletedAt,
	}

	logrus.Print("1")

	checkRecord, err := NewMySQLReopsitory().GetCutomerById(model.CustomerId)
	if err != nil {
		return nil, err
	}
	logrus.Println("Record Before changes: ", checkRecord)
	// check error ErrRecordNotFound
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}

	db.Model(&customer).Where("customer_id= ?", model.CustomerId).Updates(customerDataHolder)

	logrus.Println("Record after changes", checkRecord)

	return &customerDataHolder, err
}
