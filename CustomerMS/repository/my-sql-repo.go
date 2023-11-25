package repository

import (
	repository "CustomerMS/db"
	"CustomerMS/models"
	"CustomerMS/utils"
	"errors"

	sqlmy "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLRepository struct {
	DB *gorm.DB
}

var (
	customer          models.Customer
	dbUser            = utils.EnvVarRead(`DATABASEUSER`)
	dbPass            = utils.EnvVarRead(`DATABASEPASS`)
	connectionString  = utils.EnvVarRead(`DBCONNECTION`)
	databaseName      = utils.EnvVarRead(`DATABASE`)
	networkProtcol    = utils.EnvVarRead(`NETPROTOCOL`)
	datetimePrecision = 30
	sqlConfig         = sqlmy.Config{
		User:                 dbUser,
		Passwd:               dbPass,
		Net:                  networkProtcol,
		Addr:                 connectionString,
		DBName:               databaseName,
		ParseTime:            true,
		AllowNativePasswords: true,
	}

	gormConfig = mysql.Config{
		DSNConfig:                 &sqlConfig,         // The Required MySQL configurations
		DefaultStringSize:         256,                // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision:  true,               // disable datetime precision support, which not supported before MySQL 5.6
		DefaultDatetimePrecision:  &datetimePrecision, // default datetime precision
		DontSupportRenameIndex:    true,               // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,               // use change when rename column, rename rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,              // auto configure based on currently MySQL version
	}
)

func NewMySQLReopsitory() BillRespository {

	open, err := gorm.Open(mysql.New(gormConfig), &gorm.Config{
		FullSaveAssociations: true,
	})

	if err != nil {
		logrus.Panicf("database is unavalible: %v", err)
	}

	pingErr := repository.PingServer()
	if pingErr != nil {
		logrus.Panicf("database is currently offline: %v", err)
	}

	return &MySQLRepository{open}
}

// CreateCutomer implements BillRespository.
func (m *MySQLRepository) CreateCutomer(model *models.Customer) (*models.Customer, error) {

	customer = models.Customer{
		CustomerId:    model.CustomerId,
		CustomerName:  model.CustomerName,
		ContactNumber: model.ContactNumber,
		Address:       model.Address,
		Priority:      model.Priority,
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
		DeletedAt:     model.DeletedAt,
	}

	result := m.DB.Create(&customer)

	return &customer, result.Error

}

// GetCutomerById implements BillRespository.
func (m *MySQLRepository) GetCutomerById(id string) (models.Customer, error) {

	result := m.DB.First(&customer, "customer_id= ?", id)
	logrus.Infoln("You have enter id: " + id)
	//result := m.DB.Model(&models.Customer{CustomerId: id}).First(&customer)
	logrus.Println(result.RowsAffected)
	//logrus.Infoln(result.Statement)
	logrus.Errorln(result.Error)
	logrus.Println(result)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		logrus.Errorf("No Records Found for Entered Id: %v", id)
	}

	logrus.Println("Values "+customer.CustomerId, customer.CustomerName, customer.CreatedAt)
	return customer, result.Error

}

// GetAllCutomer implements BillRespository.
func (m *MySQLRepository) GetAllCutomer() ([]*models.Customer, error) {
	var customers []*models.Customer

	// Get all records
	result := m.DB.Find(&customers) // SELECT * FROM customers;

	logrus.Printf("No of Rows affected: %d", result.RowsAffected) // returns found records count, equals `len(users)`

	// check error ErrRecordNotFound
	errors.Is(result.Error, gorm.ErrRecordNotFound)

	return customers, result.Error
}

// This is for updateing record with Model as input params
// UpdateCutomer implements BillRespository.
func (m *MySQLRepository) UpdateCutomer(model *models.Customer) (*models.Customer, error) {

	customerDataHolder := models.Customer{
		CustomerId:    model.CustomerId,
		CustomerName:  model.CustomerName,
		ContactNumber: model.ContactNumber,
		Address:       model.Address,
		Priority:      model.Priority,
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
		DeletedAt:     model.DeletedAt,
	}

	checkRecord, err := m.GetCutomerById(model.CustomerId)
	if err != nil {
		return nil, err
	}
	logrus.Println("Record Before changes: ", checkRecord)
	// check error ErrRecordNotFound
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}

	result := m.DB.Model(&customer).Where("customer_id = ?", model.CustomerId).Updates(customerDataHolder)

	logrus.Println("Record after changes", checkRecord)

	return &customerDataHolder, result.Error
}

// DeleteCutomer implements BillRespository.
func (m *MySQLRepository) DeleteCutomer(id string) (*models.Customer, error) {

	logrus.Println("You have entered Id: " + id)
	cust, err := m.GetCutomerById(id)
	if err != nil {
		logrus.Errorf("No Records Found for Entered Id: %v", id)
		return nil, err
	}
	result := m.DB.Delete(&customer, "customer_id = ?", id)
	logrus.Println(result.RowsAffected)
	logrus.Errorln(result.Error)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		logrus.Errorf("No Records Found for Entered Id: %v", id)
	}

	return &cust, result.Error

}

func (m *MySQLRepository) Close() {
	sqlDB, err := m.DB.DB()
	if err != nil {
		logrus.Panic("Cannot close db")
	}
	sqlDB.Close()

}
