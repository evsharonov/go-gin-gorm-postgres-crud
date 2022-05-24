package config

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const projectDirName = "go-gin-gorm-postgres-crud"

type DSLConfig struct {
	Connection string
	Host       string
	Port       string
	Username   string
	Password   string
	Name       string
	Charset    string
	Sslmode    string
}

type DatabaseConfig struct {
	DB *DSLConfig
}

//Auth0.com implementation
type Auth0APIConfig struct {
	Domain   string
	Audience string
}

type Auth0Config struct {
	Auth *Auth0APIConfig
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

func GetAuth0Config() *Auth0Config {
	loadEnv()

	Domain := os.Getenv("AUTH0_DOMAIN")
	Audience := os.Getenv("AUTH0_AUDIENCE")

	return &Auth0Config{
		&Auth0APIConfig{
			Domain:   Domain,
			Audience: Audience,
		},
	}
}

func GetDBConfig() *DatabaseConfig {
	loadEnv()

	dbConnection := os.Getenv("DB_CONNECTION")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbCharset := os.Getenv("DB_CHARSET")
	dbSslmode := os.Getenv("DB_SSLMODE")

	return &DatabaseConfig{
		DB: &DSLConfig{
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
