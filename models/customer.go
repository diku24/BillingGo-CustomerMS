package models

import "github.com/jinzhu/gorm"

// Customer represents the customer for this billing application
//
// A customer is the security principal for this application.
// It's also used as one of main axes for reporting.
//
// A customer is will have all neccessary idenitity information
//
//swagger:model customer
type Customer struct {
	gorm.Model

	//the customerId is Id for the Customer.
	//
	//required: true
	//min: 1
	//unique: true
	CustomerId string `json:"customer_id" gorm:"primary_key"`

	//the  customerName is Name of the Customer.
	//required: true
	CustomerName string `json:"customer_name"`

	//the ContanctNumber is the Phone Number or mobile number of Customer.
	//required: true
	//min lenght: 10
	ContactNumber string `json:"contact_number"`

	//the address is either full address or City name of the Customer.
	//required true
	Address string `json:"address"`

	//the Priority is used to address the preference to customers request.
	Priority string `json:"priority"`
}
