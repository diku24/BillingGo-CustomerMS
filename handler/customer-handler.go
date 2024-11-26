package handler

import (
	billerr "CustomerMS/errors"
	"CustomerMS/models"
	"CustomerMS/services"

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

// Get Customers
//
//	@Summary		Get customers
//	@Description	Get all customers if no customer_is provided else get customer of the passed Id.
//	@Tags			Customers
//	@Accept			json
//	@Produce		json
//	@Param			id						path		int	true	"customer_id"
//	@Success		200						{object}	models.Customer
//	@Failure		400						{object}	errors.HTTPError
//	@Failure		404						{object}	errors.HTTPError
//	@Failure		500						{object}	errors.HTTPError
//	@Router			/customer/{customer_id}	[get]
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

// Create Customer
//
//	@Summary		Create customers
//	@Description	Create customers.
//	@Tags			Customers
//	@Accept			json
//	@Produce		json
//	@Param			customer	body		models.Customer	true	"Create customer"
//	@Success		200			{object}	models.Customer
//	@Failure		400			{object}	errors.HTTPError
//	@Failure		404			{object}	errors.HTTPError
//	@Failure		500			{object}	errors.HTTPError
//	@Router			/customer	[post]
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

// UPDATE Customers
//
//	@Summary		update customers
//	@Description	update customers.
//	@Tags			Customers
//	@Accept			json
//	@Produce		json
//	@Param			id				path		int				true	"customer_id"
//	@Param			customer		body		models.Customer	true	"Update customer"
//	@Success		200				{object}	models.Customer
//	@Failure		400				{object}	errors.HTTPError
//	@Failure		404				{object}	errors.HTTPError
//	@Failure		500				{object}	errors.HTTPError
//	@Router			/customer/{id}	[put]
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

// DELETE Cusomters
//
//	@Summary		Delete an customer
//	@Description	Delete by customer ID
//	@Tags			Customer
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Customer ID"	Format(int64)
//	@Success		204	{object}	models.Customer
//	@Failure		400	{object}	errors.HTTPError
//	@Failure		404	{object}	errors.HTTPError
//	@Failure		500	{object}	errors.HTTPError
//	@Router			/customer/{customer_id} [delete]
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
