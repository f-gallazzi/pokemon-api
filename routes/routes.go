package routes

import (
	"pokemons-challenge/services"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {

	r := gin.Default()

	services.SetupRoutesPokemon(r.Group("/pokemons"))
	services.SetupRoutesImage(r.Group("/images"))

	return r
}
