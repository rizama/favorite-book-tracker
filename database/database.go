package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBinstance struct {
	DB *gorm.DB
}

var DB DBinstance

func ConnectDB() {
	host := os.Getenv("HOST_PG")
	port := os.Getenv("PORT_PG")
	dbname := os.Getenv("DATABASE_PG")
	username := os.Getenv("USERNAME_PG")
	password := os.Getenv("PASSWORD_PG")

	dsn := "host=" + host + " user=" + username + " password='" + password + "' dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	// db.AutoMigrate(&models.Book{})

	DB = DBinstance{
		DB: db,
	}

}
