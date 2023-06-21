package model

import (
	"BillingGo/pkg/configuration"
	"time"

	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	Userid    string `gorm:"primaryKey" json:"userid"`
	Username  string `json:"username"`
	Password  string `json:"userpass"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
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
	return Users
}

func Getuserbyid(Id int64) (*User, *gorm.DB) {
	var getuser User
	db := db.Where("ID= ?", Id).Find(&getuser)
	return &getuser, db
}
func Deleteuser(ID int64) User {
	var user User
	db.Where("ID=?", ID).Delete(user)
	return user
}
