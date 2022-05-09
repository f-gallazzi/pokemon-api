package main

import (
	"log"
	"pokemons-challenge/db"
	"pokemons-challenge/routes"

	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	db.Init()
	db.AutoMigrate()
	db.InitialData()
	r := routes.Setup()
	//Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
