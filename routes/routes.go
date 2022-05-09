package routes

import (
	"pokemons-challenge/services"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {

	r := gin.Default()

	r.Use(CORSMiddleware())

	services.SetupRoutesPokemon(r.Group("/pokemons"))
	services.SetupRoutesImage(r.Group("/images"))

	return r
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
