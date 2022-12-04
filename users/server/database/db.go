package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func Connect() {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"root",
		"Osiz@123",
		"127.0.0.1:3306",
		"grpcJWT",
	)

	Instance, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
}

// func Migrate() {
// 	Instance.AutoMigrate(&respository.User{})
// 	log.Println("Database Migration Completed...")
// }
