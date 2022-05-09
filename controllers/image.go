package controllers

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"pokemons-challenge/db"
	"pokemons-challenge/models"

	"github.com/google/uuid"
)

func UploadImage(file multipart.File) *models.Image {

	filename := uuid.New()

	out, err := os.Create("public/" + filename.String())
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	image := models.Image{
		Path: filename.String(),
	}

	db.GetInstace().Create(&image)

	return &image
}

func GetByIdImage(id int) *models.Image {
	var image = models.Image{
		BaseEntity: models.BaseEntity{
			ID: id,
		},
	}

	if result := db.GetInstace().First(&image); result.Error != nil {
		return nil
	}

	return &image

}
