package db

import (
	"log"

	"github.com/evsharonov/go-gin-gorm-crud/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := "host=localhost user=postgres password=0000 dbname=GoCRUD port=5432 sslmode=disable"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	gormDB.AutoMigrate(&models.User{}, &models.Message{})

	return gormDB
}
