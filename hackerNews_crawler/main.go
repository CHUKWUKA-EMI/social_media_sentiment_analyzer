package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chukwuka-emi/hackerNews_comments_sentiments/hackerNews_crawler/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}
	router := gin.Default()
	router.SetTrustedProxies(nil)

	routes.InitRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5001"
	}

	router.Run(fmt.Sprintf(":%s", port))
}
