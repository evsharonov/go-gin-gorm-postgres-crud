package config

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const projectDirName = "go-gin-gorm-postgres-crud"

type DBConfig struct {
	Connection string
	Host       string
	Port       string
	Username   string
	Password   string
	Name       string
	Charset    string
	Sslmode    string
}

type Config struct {
	DB *DBConfig
}

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory)) //for the test using purpose

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func GetConfig() *Config {
	loadEnv()

	dbConnection := os.Getenv("DB_CONNECTION")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbCharset := os.Getenv("DB_CHARSET")
	dbSslmode := os.Getenv("DB_SSLMODE")

	return &Config{
		DB: &DBConfig{
			Connection: dbConnection,
			Host:       dbHost,
			Port:       dbPort,
			Username:   dbUsername,
			Password:   dbPassword,
			Name:       dbName,
			Charset:    dbCharset,
			Sslmode:    dbSslmode,
		},
	}
}
