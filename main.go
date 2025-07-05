package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/resetcentral/media_library/handlers"
	"github.com/resetcentral/media_library/storage"
)

func main() {
	err := storage.NewStorage()
	if err != nil {
		log.Fatalln(err)
	}

	router := gin.Default()
	router.GET("/artist", handlers.GetAllArtists)
	router.GET("/artist/:id", handlers.GetArtistByID)
	router.POST("/artist", handlers.PostArtist)
	router.DELETE("/artist/:id", handlers.DeleteArtist)
	router.Run("localhost:8000")
}
