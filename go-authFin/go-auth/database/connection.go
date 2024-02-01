package database

import (
	"github.com/ObserverVinc/Katalog_pusri/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:password@tcp(127.0.0.1:3307)/pusri-go"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = connection
	connection.AutoMigrate(&models.User{})
}
