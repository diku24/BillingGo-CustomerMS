package services_test

import (
	"BillingGo/mocks"
	"BillingGo/models"
	"BillingGo/services"
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
	control := gomock.NewController(t)

	defer control.Finish()

	mockRepo := mocks.NewMockBillRespository(control)

	mockRepo.EXPECT().CreateCutomer(gomock.Any()).Return(&mockCustomer, nil)

	testService := services.NewCustomerService(mockRepo)

	result, err := testService.Create(&mockCustomer)

	assert.Equal(t, "1", result.CustomerId)
	assert.Equal(t, "Diku", result.CustomerName)
	assert.Equal(t, "345-345-345", result.ContactNumber)
	assert.Equal(t, "Ahmednagar", result.Address)
	assert.Equal(t, "High", result.Priority)

	assert.Nil(t, err)

}

func TestGetService(t *testing.T) {

	control := gomock.NewController(t)

	defer control.Finish()

	mockRepo := mocks.NewMockBillRespository(control)

	mockRepo.EXPECT().GetAllCutomer().Return([]*models.Customer{&mockCustomer}, nil)

	testService := services.NewCustomerService(mockRepo)

	result, err := testService.GetAll()

	assert.Equal(t, "1", result[0].CustomerId)
	assert.Equal(t, "Diku", result[0].CustomerName)
	assert.Equal(t, "345-345-345", result[0].ContactNumber)
	assert.Equal(t, "Ahmednagar", result[0].Address)
	assert.Equal(t, "High", result[0].Priority)

	assert.Nil(t, err)
}

func TestGetByIdService(t *testing.T) {
	control := gomock.NewController(t)

	defer control.Finish()

	mockRepo := mocks.NewMockBillRespository(control)

	mockRepo.EXPECT().GetCutomerById(mockCustomer.CustomerId).Return(mockCustomer, nil)

	testService := services.NewCustomerService(mockRepo)

	result, err := testService.GetById(mockCustomer.CustomerId)

	assert.Equal(t, "1", result.CustomerId)
	assert.Equal(t, "Diku", result.CustomerName)
	assert.Equal(t, "345-345-345", result.ContactNumber)
	assert.Equal(t, "Ahmednagar", result.Address)
	assert.Equal(t, "High", result.Priority)

	assert.Nil(t, err)
}

func TestDeleteService(t *testing.T) {
	control := gomock.NewController(t)

	defer control.Finish()

	mockRepo := mocks.NewMockBillRespository(control)

	mockRepo.EXPECT().DeleteCutomer(mockCustomer.CustomerId).Return(nil)
	testService := services.NewCustomerService(mockRepo)
	err := testService.Delete(mockCustomer.CustomerId)
	assert.Nil(t, err)
}

func TestUpdateService(t *testing.T) {
	control := gomock.NewController(t)

	defer control.Finish()

	mockRepo := mocks.NewMockBillRespository(control)

	mockRepo.EXPECT().UpdateCutomer(&mockCustomer).Return(&mockCustomer, nil)

	testService := services.NewCustomerService(mockRepo)

	result, err := testService.Update(&mockCustomer)

	assert.Equal(t, "1", result.CustomerId)
	assert.Equal(t, "Diku", result.CustomerName)
	assert.Equal(t, "345-345-345", result.ContactNumber)
	assert.Equal(t, "Ahmednagar", result.Address)
	assert.Equal(t, "High", result.Priority)

	assert.Nil(t, err)
}
