package handler

import (
	"BillingGo/errors"
	"BillingGo/models"
	"BillingGo/services"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type CustomerController struct{}

var customerService services.BillService

func NewCustomerController(service services.BillService) BillHandler {
	customerService = service
	return &CustomerController{}
}

// DELETE implements BillHandler.
func (*CustomerController) DELETE(response http.ResponseWriter, req *http.Request) {
	panic("unimplemented")
}

// GET implements BillHandler.
func (*CustomerController) GET(response http.ResponseWriter, req *http.Request) {
	//panic("unimplemented")
	response.Header().Set("content-type", "application/json")
	//var customeri models.Customer

	customerIdParam := req.URL.Query().Get("customer_id")
	logrus.Infoln(customerIdParam)

	customer, err := customerService.GetById(customerIdParam)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		logrus.Errorln(err.Error())
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error Getting the Record From Database"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(customer)

}

// POST implements BillHandler.
func (*CustomerController) POST(response http.ResponseWriter, req *http.Request) {

	response.Header().Set("content-type", "application/json")
	var customer models.Customer

	err := json.NewDecoder(req.Body).Decode(&customer)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		logrus.Errorln(err.Error())
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error UnMarshiling the Request"})
		return
	}

	result, err := customerService.Create(&customer)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error Saving the Post Data - customer Data"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}

// PUT implements BillHandler.
func (*CustomerController) PUT(response http.ResponseWriter, req *http.Request) {
	panic("unimplemented")
}
