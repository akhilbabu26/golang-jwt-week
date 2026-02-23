package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func ConnectBD(){
	dsn := "host=localhost user=postgres password=123 dbname=jwt_project port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		panic("faild to connect database")
	}

	DB = db
}