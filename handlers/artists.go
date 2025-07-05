package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/resetcentral/media_library/models"
	"github.com/resetcentral/media_library/storage"
)

func GetAllArtists(c *gin.Context) {
	artists, err := storage.DB.FindArtists()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to load artists"})
		return
	}

	c.IndentedJSON(http.StatusOK, artists)
}

func GetArtistByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	artist, err := storage.DB.FindArtistByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No artist found with that ID"})
		return
	}

	c.IndentedJSON(http.StatusOK, artist)
}

func PostArtist(c *gin.Context) {
	var artist models.Artist
	err := c.BindJSON(&artist)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid artist data"})
		return
	}

	err = storage.DB.CreateArtist(&artist)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to save new artist"})
		return
	}

	c.IndentedJSON(http.StatusOK, artist)
}

func DeleteArtist(c *gin.Context) {
	var artist models.Artist

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	artist.ID = uint(id)
	err = storage.DB.DeleteArtist(artist)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete artist"})
	}
}
