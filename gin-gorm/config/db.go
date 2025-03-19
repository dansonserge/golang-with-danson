package config

import (
	"gin-gorm/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "postgres://postgres:1q2w3e4r@gin_gorm_db:5432/gin_gorm_db?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})
	DB = db

}
