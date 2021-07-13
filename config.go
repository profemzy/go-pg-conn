package main

import (
	"log"
	"os"
)

func goEnvVariable(key string) string {
	// load .env file
	//err := godotenv.Load(".env")

	result := os.Getenv(key)

	if result == "" {
		log.Fatalln("Error loading fetching env variable")
	}

	return result
}