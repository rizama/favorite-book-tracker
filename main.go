package main

import (
	"embed"
	"fmt"
	"log"

	httpLib "net/http"
	"os"

	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"github.com/rizama/favorite-book-tracker/delivery/http"
	"github.com/rizama/favorite-book-tracker/domain"
)

//go:embed resource/*
//go:embed resource/img/*.png
var resourcefs embed.FS

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed load env")
	}

	domain := domain.ConstructDomain()
	engine := html.NewFileSystem(httpLib.FS(resourcefs), ".html")
	app := http.NewHttpDelivery(domain, engine)
	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT_APP")))
}
