package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    Username string `gorm:"uniqueIndex;not null"`
    Password string `gorm:"not null"`
    Balance  float64
}

var DB *gorm.DB

func InitDB() error{
	var err error
	DB, err = gorm.Open(sqlite.Open("orderbook.db"), &gorm.Config{})
	if err != nil{
		return err
	}
	return DB.AutoMigrate(&User{})
}