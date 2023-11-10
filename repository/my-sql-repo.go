package repository

import (
	"BillingGo/models"
	"BillingGo/utils"
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type MySQLRepository struct{}

var (
	db               *sql.DB
	err              error
	dbUser           = utils.EnvVarRead("DATABASEUSER")
	dbPass           = utils.EnvVarRead("DATABASEPASS")
	connectionString = utils.EnvVarRead("DBCONNECTION")
	databaseName     = utils.EnvVarRead("DATABASE")
	networkProtcol   = utils.EnvVarRead("NETPROTOCOL")
	customer         models.Customer
	sqlConfig        = mysql.Config{
		User:   dbUser,
		Passwd: dbPass,
		Net:    networkProtcol,
		Addr:   connectionString,
		DBName: databaseName,
	}
)

func NewMySQLReopsitory() BillRespository {
	return &MySQLRepository{}
}

// CreateCutomer implements BillRespository.
func (*MySQLRepository) CreateCutomer(model *models.Customer) (*models.Customer, error) {
	//panic("unimplemented")
	db, err := gorm.Open("mysql", sqlConfig.FormatDSN())

	if err != nil {
		logrus.Errorln(err.Error())
		logrus.Errorln("Failed to connect to the Database!!")
	}

	customerDataHolder := models.Customer{
		CustomerId:    model.CustomerId,
		CustomerName:  model.CustomerName,
		ContactNumber: model.ContactNumber,
		Address:       model.Address,
		Priority:      model.Priority,
	}

	db.Create(&customerDataHolder)
	defer db.Close()
	return &customerDataHolder, err

}

// GetCutomerById implements BillRespository.
func (*MySQLRepository) GetCutomerById(id string) (models.Customer, error) {
	//panic("unimplemented")
	db, err := gorm.Open("mysql", sqlConfig.FormatDSN())
	if err != nil {
		logrus.Errorln(err.Error())
		logrus.Errorln("Failed to Connect to the Database !!")
	}
	logrus.Println("You have entered Id: " + id)
	result := db.First(&customer, "customer_id = ?", id)
	logrus.Println(result.RowsAffected)
	logrus.Errorln(result.Error)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		logrus.Errorf("No Records Found for Entered Id: %v", id)
	}

	logrus.Println("Values "+customer.CustomerId, customer.CustomerName)
	return customer, nil

}

// DeleteCutomer implements BillRespository.
func (*MySQLRepository) DeleteCutomer(id string) error {
	panic("unimplemented")
}

// GetAllCutomer implements BillRespository.
func (*MySQLRepository) GetAllCutomer() ([]models.Customer, error) {
	panic("unimplemented")
}

// UpdateCutomer implements BillRespository.
func (*MySQLRepository) UpdateCutomer(id string) error {
	panic("unimplemented")
}

// PingServer implements BillRespository.
// Capture connection properties.
func (*MySQLRepository) PingServer() {
	// Get a database handle.
	db, err = sql.Open("mysql", sqlConfig.FormatDSN())
	if err != nil {
		logrus.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		logrus.Fatal(pingErr)
	}
	logrus.Infoln("You are Connected to the  MySQL server " + connectionString + " Database: " + databaseName)
}
