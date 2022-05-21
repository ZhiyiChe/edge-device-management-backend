package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func init() {
	var err error
	dsn := "root:yourpassword@tcp(127.0.0.1:3306)/edge-device-management-backend?charset=utf8mb4&parseTime=True&loc=Local" // Data Source Name
	Database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("gorm.Open() failed: %v \n", err)
	}
}

type User struct {
	Id          int `gorm:"column:id;primaryKey"`
	Account     string
	Password    string
	IsAdmin     int
	PhoneNumber string
	Email       string
	LastLogin   string
}

type Apilog struct {
	Id       int `gorm:"column:id;primaryKey"`
	Url      string
	Header   string
	Body     string
	Caller   string
	CallTime string
}
