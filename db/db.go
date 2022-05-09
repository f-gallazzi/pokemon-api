package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"pokemons-challenge/models"
)

var instance *gorm.DB

func GetInstace() *gorm.DB {
	return instance
}

func Init() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	// default values
	if len(host) == 0 {
		host = "localhost"
	}
	if len(port) == 0 {
		port = "3306"
	}

	if instance == nil {
		dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?parseTime=true"
		fmt.Print("dsn: ", dsn)
		db, errDb := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if errDb != nil {
			fmt.Print("Error Database connection:", errDb)
		} else {
			instance = db
		}
	}
}

func AutoMigrate() {
	GetInstace().Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.PokemonExpansion{},
		&models.PokemonType{},
		&models.Image{},
		&models.Pokemon{},
	)
}

func InitialData() {
	//GetInstace().Exec("DELETE FROM pokemon_types")
	pokemonTypes := []models.PokemonType{
		{Name: "Agua"},
		{Name: "Fuego"},
		{Name: "Hierba"},
		{Name: "Eléctrico"},
	}
	for _, e := range pokemonTypes {
		if result := GetInstace().Where(&models.PokemonType{Name: e.Name}).First(&e); result.Error != nil {
			GetInstace().Create(&e)
		}
	}

	//GetInstace().Exec("DELETE FROM pokemon_expansions")
	pokemonExpansions := []models.PokemonExpansion{
		{Name: "Base Set"},
		{Name: "Jungle"},
		{Name: "Fossil"},
		{Name: "Base Set 2"},
	}
	for _, e := range pokemonExpansions {
		if result := GetInstace().Where(&models.PokemonExpansion{Name: e.Name}).First(&e); result.Error != nil {
			GetInstace().Create(&e)
		}
	}
}
