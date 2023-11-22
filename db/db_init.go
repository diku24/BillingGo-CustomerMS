package repository

import (
	"BillingGo/models"
	"BillingGo/utils"

	"github.com/DATA-DOG/go-sqlmock"
	sqlmy "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbUser            = utils.EnvVarRead(`DATABASEUSER`)
	dbPass            = utils.EnvVarRead(`DATABASEPASS`)
	connectionString  = utils.EnvVarRead(`DBCONNECTION`)
	databaseName      = utils.EnvVarRead(`DATABASE`)
	networkProtcol    = utils.EnvVarRead(`NETPROTOCOL`)
	customer          models.Customer
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

func OpenMysqlConnection() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(gormConfig), &gorm.Config{
		FullSaveAssociations: true,
	})
	if err != nil {
		logrus.Errorln("Error while database connection")
		return nil, err
	}
	return db, err
}

// Capture connection properties.
func PingServer() {
	// Get a database handle.
	db, err := OpenMysqlConnection()
	if err != nil {
		logrus.Fatal(err)
	}

	//DB.DB() returns the pointer which points to the SQL database server we opened and an error value
	sql, err := db.DB()
	//If Db.DB() returns non nil error, it means the database was unreachable, facing internal errors, or improperly configured etc.
	if err != nil {
		panic(err.Error())
	}

	pingErr := sql.Ping()
	if pingErr != nil {
		logrus.Fatal(pingErr)
	}
	logrus.Infoln("You are Connected to the  MySQL server " + connectionString + " Database: " + databaseName)
}

// To create the database if not exists
func DatabaseCreation() error {
	return nil
}

func TableCreation() error {
	db, err := OpenMysqlConnection()
	if err != nil {
		return err
	}

	if !db.Migrator().HasTable(&customer) {
		//Drop Customer Table if it is present alredy
		droptablerr := db.Migrator().DropTable(&customer)
		if droptablerr != nil {
			return droptablerr
		}
		logrus.Println("Table Dropped Successfully")

		//create customer Table
		createtablerr := db.AutoMigrate(&customer)
		if createtablerr != nil {
			return createtablerr
		}
		logrus.Println("Table Creates Successfully")
	}

	return nil

}

func NewMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()

	if err != nil {
		logrus.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if err != nil {
		logrus.Fatalf("An error '%s' was not expected when opening gorm database", err)
	}

	return gormDB, mock
}
