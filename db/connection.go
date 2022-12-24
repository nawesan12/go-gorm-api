package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DSN string = " "
var DB *gorm.DB

func DBConnection() {
	var err error
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Database connected")
}
