package main

import (
	"log"

	server "github.com/chukwuka-emi/hackerNews_comments_sentiments/sentiment_analyzer/router"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}

	server := server.Server{}
	server.Start()
}
