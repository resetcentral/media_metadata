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
	router.GET("/artist", handlers.GetArtists)
	router.GET("/artist/:id", handlers.GetArtistByID)
	router.POST("/artist", handlers.PostArtist)
	router.DELETE("/artist/:id", handlers.DeleteArtist)

	router.GET("/studio", handlers.GetStudios)
	router.GET("/studio/:id", handlers.GetStudioByID)
	router.POST("/studio", handlers.PostStudio)
	router.DELETE("/studio/:id", handlers.DeleteStudio)

	router.GET("/tag", handlers.GetTags)
	router.GET("/tag/:id", handlers.GetTagByID)
	router.POST("/tag", handlers.PostTag)
	router.DELETE("/tag/:id", handlers.DeleteTag)

	router.Run("localhost:8000")
}
