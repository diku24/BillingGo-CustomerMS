package model

import (
	"BillingGo/pkg/configuration"
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	Userid    string         `gorm:"primaryKey" json:"userid"`
	Username  string         `json:"username"`
	Password  string         `json:"userpass"`
	CreatedAt time.Time      `json:"createdtime"`
	UpdatedAt time.Time      `json:"updatedtime"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedtime"`
}

func init() {
	configuration.Connect()
	db = configuration.GetDB()
	db.AutoMigrate(&User{})
}

func (b *User) Createuser() *User {
	db.FirstOrCreate(b)
	db.Create(&b)
	return b
}
func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	if len(Users) == 0 {
		log.Printf("Record not found in database: %s", gorm.ErrRecordNotFound)
	}
	return Users
}

func Getbyid(Id string) (User, *gorm.DB, error) {
	var getuser User
	db := db.Where("userid= ?", Id).First(&getuser)
	if db.Error != nil {
		errors.Is(db.Error, gorm.ErrRecordNotFound)
		log.Printf("Record not found for userId: %s", Id)
	}
	return getuser, db, db.Error
}

func Deleteuser(userid string) User {
	var user User
	deletedUser, _, _ := Getbyid(userid)
	db.Where("userid=?", userid).Delete(&user)
	return deletedUser
}
