package repository_test

import (
	db "BillingGo/db"
	"BillingGo/models"
	"BillingGo/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var (
	mockCustomer = models.Customer{
		CustomerId:    "1",
		CustomerName:  "Diku",
		ContactNumber: "345-345-345",
		Address:       "Ahmednagar",
		Priority:      "high",
	}
)

func TestGetById(t *testing.T) {

	db, mock := db.NewMockDB()

	rows := sqlmock.NewRows([]string{"CustomerId", "CustomerName", "ContactNumber", "Address", "Priority"}).AddRow(mockCustomer.CustomerId, mockCustomer.CustomerName, mockCustomer.ContactNumber, mockCustomer.Address, mockCustomer.Priority)

	//query := `SELECT\s.*FROM\s.*WHERE\s.*`
	query := `SELECT\s.*FROM\s.customers. WHERE customer_id= \\?`
	mock.ExpectQuery(query).WillReturnRows(rows)

	var customer models.Customer
	logrus.Println("Customer at cretion time: " + customer.CustomerId + customer.Priority + customer.CustomerName)

	err := db.First(&customer, "customer_id= ?", mockCustomer.CustomerId).Error
	if err != nil {
		t.Fatalf("Error in finding the Customer: %v", err)
	}
	logrus.Println(customer.CustomerId)
	logrus.Println(customer.CustomerName)
	logrus.Println(customer.ContactNumber)
	logrus.Println(customer.Address)
	logrus.Println(customer.Priority)

	custinfo, err := repository.NewMySQLReopsitory().GetCutomerById(mockCustomer.CustomerId)

	assert.NotNil(t, custinfo)
	assert.NoError(t, err)
	assert.Equal(t, mockCustomer.CustomerId, customer.CustomerId)
	assert.Equal(t, mockCustomer.CustomerName, customer.CustomerName)
	assert.Equal(t, mockCustomer.ContactNumber, customer.ContactNumber)
	assert.Equal(t, mockCustomer.Address, customer.Address)
	assert.Equal(t, mockCustomer.Priority, customer.Priority)
	assert.Nil(t, err)
}

func TestGetAll(t *testing.T) {

	db, mock := db.NewMockDB()

	rows := sqlmock.NewRows([]string{"CustomerId", "CustomerName", "ContactNumber", "Address", "Priority"}).AddRow(mockCustomer.CustomerId, mockCustomer.CustomerName, mockCustomer.ContactNumber, mockCustomer.Address, mockCustomer.Priority)

	//mock.ExpectQuery(`SELECT\s.* FROM\s.*customers\s.* WHERE\s.* customer_id= ?\s.*`).WillReturnRows(rows)
	mock.ExpectQuery(`SELECT\s.*FROM\s.*WHERE\s.*`).WillReturnRows(rows)

	var customers []models.Customer

	if err := db.Find(&customers).Error; err != nil {
		t.Fatalf("Error in finding the Customers: %v", err)
	}

	if len(customers) != 1 || customers[0].CustomerName != "John Doe" {
		t.Fatalf("Unexpected user data retrieved: %v", customers)
	}

}
