package services

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"

	"pokemons-challenge/controllers"

	"pokemons-challenge/models"
)

func SetupRoutesPokemon(r *gin.RouterGroup) {
	r.GET(".", func(c *gin.Context) {
		// get all service
		c.JSON(200, controllers.GetAllPokemon())
	})

	r.GET("/:id", func(c *gin.Context) {
		// get by id service
		ID, errParam := strconv.Atoi(c.Param("id"))

		if errParam != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
			return
		}

		result := controllers.GetByIdPokemon(ID)

		if result == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
			return
		}

		c.JSON(200, result)
	})

	r.POST(".", func(c *gin.Context) {
		// create a new entity service
		var data models.Pokemon
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := controllers.CreatePokemon(&data)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, result)
	})

	r.PATCH("/:id", func(c *gin.Context) {
		// update entity by id service
		var data models.Pokemon
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ID, errParam := strconv.Atoi(c.Param("id"))

		if errParam != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
			return
		}

		result, err := controllers.UpdateByIdPokemon(ID, &data)

		if result == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
			return
		}

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, result)

	})

	r.DELETE("/:id", func(c *gin.Context) {
		// delete entity by id service
		ID, errParam := strconv.Atoi(c.Param("id"))

		if errParam != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
			return
		}

		result := controllers.DeleteByIdPokemon(ID)

		if result == false {
			c.JSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
			return
		}

		c.Done()

	})

	r.DELETE("/all", func(c *gin.Context) {
		// delete all service

		controllers.DeleteAllPokemon()

		c.Done()

	})
}
