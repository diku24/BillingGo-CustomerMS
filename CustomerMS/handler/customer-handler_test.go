package handler_test

import (
	"CustomerMS/handler"
	"CustomerMS/mocks"
	"CustomerMS/models"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var (
	testResults []models.Customer
	testResult  models.Customer
	// mockCustomer = models.Customer{
	// 	CustomerId:    "1",
	// 	CustomerName:  "Diku",
	// 	ContactNumber: "345-345-345",
	// 	Address:       "Ahmednagar",
	// 	Priority:      "High",
	// }
	mockFailCustomer = models.Customer{
		CustomerId:    "",
		CustomerName:  "Diku",
		ContactNumber: "345-345-345",
		Address:       "Ahmednagar",
		Priority:      "High",
	}

	mockCustomerBytes = []byte(`{
		"customer_id":    "1",
		"customer_name":  "Diku",
		"contact_number": "345-345-345",
		"address":       "Ahmednagar",
		"priority":      "High"
		}`)
)

const (
	uriCustomer = "/customer"
)

func generateMockService(t *testing.T) *mocks.MockBillService {
	control := gomock.NewController(t)
	defer control.Finish()
	mockService := mocks.NewMockBillService(control)

	return mockService

}

func TestPOST(t *testing.T) {

	mockCustomer := models.Customer{
		CustomerId:    "1",
		CustomerName:  "Diku",
		ContactNumber: "345-345-345",
		Address:       "Ahmednagar",
		Priority:      "High",
	}

	mockService := generateMockService(t)

	mockService.EXPECT().Create(&mockCustomer).Return(&mockCustomer, nil)

	testControl := handler.NewCustomerController(mockService)

	response := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodPost, uriCustomer, bytes.NewBuffer(mockCustomerBytes))

	err := testControl.POST(response, request)
	if err != nil {
		t.Errorf(err.Error())
	}

	result := response.Result()

	//Decode the http responce
	errr := json.NewDecoder(io.Reader(response.Body)).Decode(&testResult)
	if errr != nil {
		t.Error(errr.Error())
	}

	//assertions for failing
	assert.Equal(t, http.StatusOK, result.StatusCode)
	assert.Equal(t, mockCustomer.CustomerId, testResult.CustomerId)
	assert.Equal(t, mockCustomer.CustomerName, testResult.CustomerName)
	assert.Equal(t, mockCustomer.ContactNumber, testResult.ContactNumber)
	assert.Equal(t, mockCustomer.Address, testResult.Address)
	assert.Equal(t, mockCustomer.Priority, testResult.Priority)
}

func TestPOSTFail(t *testing.T) {

	mockCustomer := models.Customer{
		CustomerId:    "1",
		CustomerName:  "Diku",
		ContactNumber: "345-345-345",
		Address:       "Ahmednagar",
		Priority:      "High",
	}

	mockService := generateMockService(t)

	//mockService.EXPECT().Create(&mockCustomer).Return(&mockCustomer, nil)
	call1 := mockService.EXPECT().Create(&mockCustomer)
	call2 := mockService.EXPECT().Create(nil)

	mockErr := errors.New("Mock err")
	call1.Return(&mockCustomer, nil)
	call2.Return(nil, mockErr)

	testControl := handler.NewCustomerController(mockService)

	response := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodPost, uriCustomer, bytes.NewBuffer(nil))

	err := testControl.POST(response, request)
	if err != nil {
		t.Errorf(err.Error())
	}

	result := response.Result()

	//Decode the http responce
	errr := json.NewDecoder(io.Reader(response.Body)).Decode(&testResult)
	if errr != nil {
		t.Error(errr.Error())
	}

	//assertions for failing
	assert.Equal(t, http.StatusBadRequest, result.StatusCode)
	//assert.Equal(t, http.StatusInternalServerError, result.StatusCode)
	//assert.Empty(t, testResult)

}

func TestGETAll(t *testing.T) {

	mockCustomer := models.Customer{
		CustomerId:    "1",
		CustomerName:  "Diku",
		ContactNumber: "345-345-345",
		Address:       "Ahmednagar",
		Priority:      "High",
	}

	mockService := generateMockService(t)

	mockService.EXPECT().GetAll().Return([]*models.Customer{&mockCustomer}, nil)

	testController := handler.NewCustomerController(mockService)

	response := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodGet, uriCustomer, nil)

	err := testController.GET(response, request)
	if err != nil {
		t.Error(err.Error())
	}

	jsonErr := json.NewDecoder(io.Reader(response.Body)).Decode(&testResults)
	if jsonErr != nil {
		t.Error(jsonErr)
	}
	result := response.Result()

	assert.Equal(t, http.StatusOK, result.StatusCode)
	assert.Equal(t, mockCustomer.CustomerId, testResults[0].CustomerId)
	assert.Equal(t, mockCustomer.CustomerName, testResults[0].CustomerName)
	assert.Equal(t, mockCustomer.ContactNumber, testResults[0].ContactNumber)
	assert.Equal(t, mockCustomer.Address, testResults[0].Address)
	assert.Equal(t, mockCustomer.Priority, testResults[0].Priority)
}

func TestGETById(t *testing.T) {

	mockCustomer := models.Customer{
		CustomerId:    "1",
		CustomerName:  "Diku",
		ContactNumber: "345-345-345",
		Address:       "Ahmednagar",
		Priority:      "High",
	}
	mockService := generateMockService(t)

	mockService.EXPECT().GetById(mockCustomer.CustomerId).Return(mockCustomer, nil)

	testController := handler.NewCustomerController(mockService)

	response := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodGet, uriCustomer+"?customer_id="+mockCustomer.CustomerId, nil)

	err := testController.GET(response, request)
	if err != nil {
		t.Error(err.Error())
	}

	jsonErr := json.NewDecoder(io.Reader(response.Body)).Decode(&testResult)
	if jsonErr != nil {
		t.Error(jsonErr)
	}
	result := response.Result()

	//assertions for failing
	assert.Equal(t, http.StatusOK, result.StatusCode)
	assert.Equal(t, mockCustomer.CustomerId, testResult.CustomerId)
	assert.Equal(t, mockCustomer.CustomerName, testResult.CustomerName)
	assert.Equal(t, mockCustomer.ContactNumber, testResult.ContactNumber)
	assert.Equal(t, mockCustomer.Address, testResult.Address)
	assert.Equal(t, mockCustomer.Priority, testResult.Priority)
}

func TestUpdate(t *testing.T) {

	mockCustomer := models.Customer{
		CustomerId:    "1",
		CustomerName:  "Diku",
		ContactNumber: "345-345-345",
		Address:       "Ahmednagar",
		Priority:      "High",
	}

	mockService := generateMockService(t)

	mockService.EXPECT().Update(&mockCustomer).Return(&mockCustomer, nil)

	testController := handler.NewCustomerController(mockService)

	response := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodPut, uriCustomer, bytes.NewBuffer(mockCustomerBytes))

	err := testController.PUT(response, request)
	if err != nil {
		t.Errorf(err.Error())
	}

	jsonErr := json.NewDecoder(io.Reader(response.Body)).Decode(&testResult)
	if jsonErr != nil {
		t.Error(jsonErr)
	}

	result := response.Result()

	//assertions for failing
	assert.Equal(t, http.StatusOK, result.StatusCode)
	assert.Equal(t, mockCustomer.CustomerId, testResult.CustomerId)
	assert.Equal(t, mockCustomer.CustomerName, testResult.CustomerName)
	assert.Equal(t, mockCustomer.ContactNumber, testResult.ContactNumber)
	assert.Equal(t, mockCustomer.Address, testResult.Address)
	assert.Equal(t, mockCustomer.Priority, testResult.Priority)

}

func TestDELETE(t *testing.T) {

	mockCustomer := models.Customer{
		CustomerId:    "1",
		CustomerName:  "Diku",
		ContactNumber: "345-345-345",
		Address:       "Ahmednagar",
		Priority:      "High",
	}
	mockService := generateMockService(t)

	//mockService.EXPECT().Delete(mockCustomer.CustomerId).Return(&mockCustomer, nil)
	mockService.EXPECT().Delete(mockCustomer.CustomerId).Return(&mockCustomer, nil)

	testController := handler.NewCustomerController(mockService)

	response := httptest.NewRecorder()

	logrus.Println("URI: " + uriCustomer + "/" + mockCustomer.CustomerId)

	request := httptest.NewRequest(http.MethodDelete, uriCustomer+"/"+mockCustomer.CustomerId, nil)
	request = mux.SetURLVars(request, map[string]string{"customer_id": mockCustomer.CustomerId})

	err := testController.DELETE(response, request)
	if err != nil {
		t.Error(err.Error())
	}

	jsonErr := json.NewDecoder(io.Reader(response.Body)).Decode(&testResult)
	if jsonErr != nil {
		t.Error(jsonErr)
	}
	result := response.Result()

	//assertions for failing
	assert.Equal(t, http.StatusOK, result.StatusCode)
	assert.Equal(t, mockCustomer.CustomerId, testResult.CustomerId)
	assert.Equal(t, mockCustomer.CustomerName, testResult.CustomerName)
	assert.Equal(t, mockCustomer.ContactNumber, testResult.ContactNumber)
	assert.Equal(t, mockCustomer.Address, testResult.Address)
	assert.Equal(t, mockCustomer.Priority, testResult.Priority)
}
