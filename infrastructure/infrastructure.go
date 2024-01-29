package infrastructure

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection() *gorm.DB {
	host := os.Getenv("HOST_PG")
	port := os.Getenv("PORT_PG")
	dbname := os.Getenv("DATABASE_PG")
	username := os.Getenv("USERNAME_PG")
	password := os.Getenv("PASSWORD_PG")

	dsn := "host=" + host + " user=" + username + " password='" + password + "' dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Asia/Jakarta"

	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return d
}
