package controllers

import (
	"errors"
	"pokemons-challenge/db"
	"pokemons-challenge/models"
	"time"
)

func validatePokemon(pokemon *models.Pokemon) (err error) {
	// validations
	if pokemon.HP < 0 || pokemon.HP%10 != 0 {
		return errors.New("The value for field 'HP' is not a multiple of 10")
	}

	if !(pokemon.Rarity == "common" || pokemon.Rarity == "no-common" || pokemon.Rarity == "rare") {
		return errors.New("Invalid rarity value [common|no-common|rare]")
	}

	return
}

func CreatePokemon(pokemon *models.Pokemon) (*models.Pokemon, error) {

	pokemon.CreatedDate = time.Now()

	if err := validatePokemon(pokemon); err != nil {
		return nil, err
	}

	db.GetInstace().Create(pokemon)

	return pokemon, nil
}

func UpdateByIdPokemon(id int, pokemon *models.Pokemon) (*models.Pokemon, error) {

	var pokemonToUpdate models.Pokemon

	pokemonToUpdate.ID = id

	if result := db.GetInstace().First(&pokemonToUpdate); result.Error != nil {
		return nil, result.Error
	}

	pokemonToUpdate.HP = pokemon.HP
	pokemonToUpdate.IsFirstEdition = pokemon.IsFirstEdition
	pokemonToUpdate.Name = pokemon.Name
	pokemonToUpdate.Price = pokemon.Price
	pokemonToUpdate.Rarity = pokemon.Rarity

	if err := validatePokemon(&pokemonToUpdate); err != nil {
		return nil, err
	}

	db.GetInstace().Save(&pokemonToUpdate)

	return &pokemonToUpdate, nil
}

func DeleteByIdPokemon(id int) bool {

	if result := db.GetInstace().Delete(&models.Pokemon{}, id); result.Error != nil {
		return false
	}
	return true
}

func DeleteAllPokemon() {
	db.GetInstace().Exec("DELETE FROM pokemons")
}

func GetByIdPokemon(id int) *models.Pokemon {
	var pokemon = models.Pokemon{
		BaseEntity: models.BaseEntity{
			ID: id,
		},
	}

	if result := db.GetInstace().
		Joins("Image").
		Joins("PokemonExpansion").
		Joins("PokemonType").First(&pokemon); result.Error != nil {
		return nil
	}

	return &pokemon

}

func GetAllPokemon() *[]models.Pokemon {

	var pokemons []models.Pokemon

	db.GetInstace().
		Joins("Image").
		Joins("PokemonExpansion").
		Joins("PokemonType").Find(&pokemons)

	return &pokemons
}
