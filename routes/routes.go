package routes

import (
	"pokemons-challenge/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {

	r := gin.Default()

	r.Use(cors.Default())

	services.SetupRoutesPokemon(r.Group("/pokemons"))
	services.SetupRoutesImage(r.Group("/images"))

	return r
}
