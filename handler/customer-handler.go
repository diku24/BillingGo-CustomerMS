package handler

import (
	"BillingGo/errors"
	"BillingGo/models"
	"BillingGo/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type CustomerController struct{}

var customerService services.BillService

const (
	contentType     = "content-type"
	applicationJson = "application/json"
)

func NewCustomerController(service services.BillService) BillHandler {
	customerService = service
	return &CustomerController{}
}

// DELETE implements BillHandler.
func (*CustomerController) DELETE(response http.ResponseWriter, req *http.Request) error {

	customerId := mux.Vars(req)["customer_id"]
	logrus.Infoln("Id to be deleted - Handler :", customerId)
	err := customerService.Delete(customerId)
	if err != nil {
		return err
	}
	return WriteJSON(response, http.StatusOK, map[string]string{"deleted customer with Id : ": customerId})

}

// GET implements BillHandler.
func (*CustomerController) GET(response http.ResponseWriter, req *http.Request) error {

	customerIdParam := mux.Vars(req)["customer_id"]
	logrus.Infoln(customerIdParam)
	if customerIdParam != "" {
		customer, err := customerService.GetById(customerIdParam)
		if err != nil {
			return WriteJSON(response, http.StatusInternalServerError, errors.ControllerError{Message: "Error Getting the Record From Database"})
		}
		return WriteJSON(response, http.StatusOK, customer)
	} else {
		customer, err := customerService.GetAll()
		if err != nil {
			return WriteJSON(response, http.StatusInternalServerError, errors.ControllerError{Message: "Error Getting the Record From Database"})
		}
		return WriteJSON(response, http.StatusOK, customer)
	}

}

// POST implements BillHandler.
func (*CustomerController) POST(response http.ResponseWriter, req *http.Request) error {

	var customer models.Customer

	err := json.NewDecoder(req.Body).Decode(&customer)
	if err != nil {
		return WriteJSON(response, http.StatusInternalServerError, errors.ControllerError{Message: "Error UnMarshiling the Request"})
	}

	result, err := customerService.Create(&customer)
	if err != nil {
		return WriteJSON(response, http.StatusInternalServerError, errors.ControllerError{Message: "Error Saving the Post Data - customer Data"})
	}
	return WriteJSON(response, http.StatusOK, result)
}

// PUT implements BillHandler.
func (*CustomerController) PUT(response http.ResponseWriter, req *http.Request) error {

	var customer models.Customer
	err := json.NewDecoder(req.Body).Decode(&customer)
	if err != nil {
		return WriteJSON(response, http.StatusInternalServerError, errors.ControllerError{Message: "Error Reading the Params"})
	}

	tempCustomer, err := customerService.Update(&customer)
	if err != nil {
		return WriteJSON(response, http.StatusInternalServerError, errors.ControllerError{Message: "Error Reading the Params"})
	}

	logrus.Infoln("Customer after Updated - handler: ", tempCustomer)
	return WriteJSON(response, http.StatusOK, map[string]string{"Update Customer with Id: ": customer.CustomerId})

}

func WriteJSON(resp http.ResponseWriter, status int, v any) error {
	resp.Header().Set(contentType, applicationJson)
	resp.WriteHeader(status)
	return json.NewEncoder(resp).Encode(v)
}
