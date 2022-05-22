package db

import (
	"fmt"
	"log"

	"github.com/evsharonov/go-gin-gorm-crud/pkg/config"
	"github.com/evsharonov/go-gin-gorm-crud/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	config := config.GetConfig().DB

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v", config.Host, config.Username, config.Password, config.Name, config.Port, config.Sslmode)
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	gormDB.AutoMigrate(&models.User{}, &models.Message{})

	return gormDB
}
