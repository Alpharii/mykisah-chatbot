package main

import (
	// "ai-chat/config"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main(){
	if err := godotenv.Load(); err!= nil {
		log.Fatal("cannot load encv")
	}

	// config.ConnectDB()

	app:= fiber.New()

	app.Get("/", func (c*fiber.Ctx) error {
		return c.JSON("hi from backend")
	})

	log.Fatal(app.Listen(":"+ os.Getenv("PORT")))
}