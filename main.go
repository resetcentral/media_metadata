package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/resetcentral/media_library/handlers"
	"github.com/resetcentral/media_library/models"
	"github.com/resetcentral/media_library/storage"
)

func main() {
	err := storage.NewStorage()
	if err != nil {
		log.Fatalln(err)
	}

	testAudio := models.AudioMedia{
		Title:         "Test Audio 2",
		Description:   "This is a test audio",
		FilePath:      "/path/to/file2.mp3",
		Studio:        models.Studio{Name: "Dummy Studio"},
		Artists:       []models.Artist{{Name: "Artist A"}, {Name: "Artist B"}},
		ReleaseDate:   time.Now(),
		Tags:          []models.Tag{{Value: "Cool"}, {Value: "Fun"}},
		Length:        time.Duration(5 * 1e9),
		Container:     "mp3",
		AudioCodec:    "aac",
		AudioChannels: 1,
	}
	err = storage.DB.CreateMedia(&testAudio)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", testAudio)

	router := gin.Default()
	router.GET("/artist", handlers.GetAllArtists)
	router.GET("/artist/:id", handlers.GetArtistByID)
	router.POST("/artist", handlers.PostArtist)
	router.DELETE("/artist/:id", handlers.DeleteArtist)
	router.Run("localhost:8000")
}
