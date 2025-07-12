package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/resetcentral/media_metadata/models"
	"github.com/resetcentral/media_metadata/storage"
)

func GetArtists(c *gin.Context) {
	var artists []models.Artist
	var err error

	name := c.Query("name")
	artists, err = storage.DB.FindArtists(name)
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
	if err != nil || artist.ID != uint(id) {
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
	if artist.Name == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "'name' field is required"})
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
