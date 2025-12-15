package main

import (
	"ai-chat/config"
	"ai-chat/internal/router"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main(){
	if err := godotenv.Load(); err!= nil {
		log.Fatal("cannot load encv")
	}

	config.InitGemini()

	db := config.ConnectDB()
	app := router.InitRouter(db)

	app.Listen(":" + os.Getenv("PORT"))
}