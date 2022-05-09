package models

import "time"

type Pokemon struct {
	BaseEntity
	Name           string  `binding:"required"`
	HP             int     `binding:"required"`
	IsFirstEdition bool    `binding:"required"`
	Rarity         string  `binding:"required"`
	Price          float32 `binding:"required"`

	PokemonTypeID int
	PokemonType   PokemonType

	PokemonExpansionID int
	PokemonExpansion   PokemonExpansion

	ImageID int
	Image   Image

	CreatedDate time.Time
}
