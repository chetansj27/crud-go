package models

import (
	"github.com/chetansj27/crud-go/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	Id      uint64 `gorm:"primaryKey" json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func AddUser(us *User) *User {
	db.Create(&us)
	return us
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(id uint64) (*User, *gorm.DB) {
	var user *User
	db.Find(&user, id)
	return user, db
}

func DeleteUserById(id uint64) *User {
	var user *User
	db.Delete(&user, id)
	return user
}
