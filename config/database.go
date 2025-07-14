package config

import(
	"fmt"
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"playground/internal/models"
)

var DB *gorm.DB

func ConnectDataBase() {
	// get connection info from env
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSL_MODE")
	
	// create connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", dbHost, dbUser, dbPassword, dbName, dbPort, dbSSLMode)

	var err error
	DB, err =  gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connected to database")

	err = DB.AutoMigrate(&models.User{}, &models.Game{}, &models.Score{})
	if err != nil{
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migrated successfully")

}