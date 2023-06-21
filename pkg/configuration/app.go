package configuration

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "root:password@tcp(localhost:3306)/billing?charset=utf8&parseTime=True&loc=Local"
	//if thair is error goes in err and if not then the vaiue goes in d which is good connection to db
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
