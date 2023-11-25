package services_test

import (
	"CustomerMS/mocks"
	"CustomerMS/models"
	"CustomerMS/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var (
	mockCustomer = models.Customer{
		CustomerId:    "1",
		CustomerName:  "Diku",
		ContactNumber: "345-345-345",
		Address:       "Ahmednagar",
		Priority:      "High",
	}
)

func TestCreateService(t *testing.T) {

	mockRepo := genrateMockRepo(t)

	mockRepo.EXPECT().CreateCutomer(gomock.Any()).Return(&mockCustomer, nil)

	testService := services.NewCustomerService(mockRepo)

	result, err := testService.Create(&mockCustomer)

	assertMe(t, mockCustomer, *result, err)

}

func TestGetService(t *testing.T) {

	mockRepo := genrateMockRepo(t)

	mockRepo.EXPECT().GetAllCutomer().Return([]*models.Customer{&mockCustomer}, nil)

	testService := services.NewCustomerService(mockRepo)

	result, err := testService.GetAll()

	assert.Equal(t, mockCustomer.CustomerId, result[0].CustomerId)
	assert.Equal(t, mockCustomer.CustomerName, result[0].CustomerName)
	assert.Equal(t, mockCustomer.ContactNumber, result[0].ContactNumber)
	assert.Equal(t, mockCustomer.Address, result[0].Address)
	assert.Equal(t, mockCustomer.Priority, result[0].Priority)

	assert.Nil(t, err)
}

func TestGetByIdService(t *testing.T) {

	mockRepo := genrateMockRepo(t)

	mockRepo.EXPECT().GetCutomerById(mockCustomer.CustomerId).Return(mockCustomer, nil)

	testService := services.NewCustomerService(mockRepo)

	result, err := testService.GetById(mockCustomer.CustomerId)

	assertMe(t, mockCustomer, result, err)
}

func TestDeleteService(t *testing.T) {

	mockRepo := genrateMockRepo(t)
	mockRepo.EXPECT().DeleteCutomer(mockCustomer.CustomerId).Return(&mockCustomer, nil)
	testService := services.NewCustomerService(mockRepo)
	result, err := testService.Delete(mockCustomer.CustomerId)

	assertMe(t, mockCustomer, *result, err)
}

func TestUpdateService(t *testing.T) {
	mockRepo := genrateMockRepo(t)

	mockRepo.EXPECT().UpdateCutomer(&mockCustomer).Return(&mockCustomer, nil)

	testService := services.NewCustomerService(mockRepo)

	result, err := testService.Update(&mockCustomer)

	assertMe(t, mockCustomer, *result, err)
}

func genrateMockRepo(t *testing.T) *mocks.MockBillRespository {
	control := gomock.NewController(t)
	defer control.Finish()
	mockRepo := mocks.NewMockBillRespository(control)
	return mockRepo
}

func assertMe(t *testing.T, mockCustomer models.Customer, result models.Customer, err error) {

	//assertions for failing
	assert.Equal(t, mockCustomer.CustomerId, result.CustomerId)
	assert.Equal(t, mockCustomer.CustomerName, result.CustomerName)
	assert.Equal(t, mockCustomer.ContactNumber, result.ContactNumber)
	assert.Equal(t, mockCustomer.Address, result.Address)
	assert.Equal(t, mockCustomer.Priority, result.Priority)
	assert.Nil(t, err)
}
