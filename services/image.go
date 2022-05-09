package services

import (
	"fmt"
	"net/http"
	"pokemons-challenge/controllers"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetupRoutesImage(r *gin.RouterGroup) {

	r.POST("/upload", func(c *gin.Context) {

		file, _, err := c.Request.FormFile("file")

		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
			return
		}

		c.JSON(200, controllers.UploadImage(file))

	})

	r.GET("/:id", func(c *gin.Context) {

		// get by id service
		ID, errParam := strconv.Atoi(c.Param("id"))

		if errParam != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
			return
		}

		result := controllers.GetByIdImage(ID)

		if result == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
			return
		}

		c.JSON(200, result)

	})

	r.StaticFS("/file", http.Dir("public"))

	r.GET("/download/:id", func(c *gin.Context) {

		// get by id service
		ID, errParam := strconv.Atoi(c.Param("id"))

		if errParam != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
			return
		}

		result := controllers.GetByIdImage(ID)

		if result == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
			return
		}

		c.File("public/" + result.Path)

	})

}
