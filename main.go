package main

import (
	"log"

	"github.com/joho/godotenv"
	"golang.clean.architecture/api"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	api.Init()
}
