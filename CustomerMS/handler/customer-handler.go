package handler

import (
	billerr "BillingGo/errors"
	"BillingGo/models"
	"BillingGo/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type CustomerController struct {
	service services.BillService
}

const (
	contentType     = "content-type"
	applicationJson = "application/json"
)

func NewCustomerController(service services.BillService) BillHandler {

	return &CustomerController{
		service: service,
	}
}

// GET implements BillHandler.
func (s *CustomerController) GET(response http.ResponseWriter, req *http.Request) error {

	customerIdParam := req.URL.Query().Get("customer_id")
	logrus.Infoln(customerIdParam)
	if customerIdParam != "" {
		customer, err := s.service.GetById(customerIdParam)
		if err != nil {
			return WriteJSON(response, http.StatusInternalServerError, billerr.ControllerError{Message: "Error Getting the Record From Database"})
		}
		return WriteJSON(response, http.StatusOK, customer)
	} else {
		customer, err := s.service.GetAll()
		if err != nil {
			return WriteJSON(response, http.StatusInternalServerError, billerr.ControllerError{Message: "Error Getting the Record From Database"})
		}
		return WriteJSON(response, http.StatusOK, customer)
	}

}

// POST implements BillHandler.
func (s *CustomerController) POST(response http.ResponseWriter, req *http.Request) error {

	var customer models.Customer

	err := json.NewDecoder(req.Body).Decode(&customer)

	if err != nil {
		return WriteJSON(response, http.StatusBadRequest, billerr.ControllerError{Message: "Error UnMarshiling the Request"})
	}

	result, err := s.service.Create(&customer)
	if err != nil {
		return WriteJSON(response, http.StatusInternalServerError, billerr.ControllerError{Message: "Error Saving the Post Data - customer Data"})
	}
	return WriteJSON(response, http.StatusOK, result)

}

// PUT implements BillHandler.
func (s *CustomerController) PUT(response http.ResponseWriter, req *http.Request) error {

	var customer models.Customer
	err := json.NewDecoder(req.Body).Decode(&customer)
	if err != nil {
		return WriteJSON(response, http.StatusInternalServerError, billerr.ControllerError{Message: "Error Reading the Params"})
	}

	tempCustomer, err := s.service.Update(&customer)
	if err != nil {
		return WriteJSON(response, http.StatusInternalServerError, billerr.ControllerError{Message: "Error Reading the Params"})
	}

	logrus.Infoln("Customer after Updated - handler: ", tempCustomer)
	return WriteJSON(response, http.StatusOK, tempCustomer)

}

// DELETE implements BillHandler.
func (s *CustomerController) DELETE(response http.ResponseWriter, req *http.Request) error {

	customerId := mux.Vars(req)["customer_id"]
	//customerId := req.URL.Query().Get("customer_id")
	logrus.Infoln("Id to be deleted - Handler : ", customerId)
	resultCustomer, err := s.service.Delete(customerId)
	if err != nil {
		return err
	}
	return WriteJSON(response, http.StatusOK, resultCustomer)

}

func WriteJSON(resp http.ResponseWriter, status int, v any) error {
	resp.Header().Set(contentType, applicationJson)
	resp.WriteHeader(status)
	return json.NewEncoder(resp).Encode(v)
}
