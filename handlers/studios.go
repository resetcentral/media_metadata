package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/resetcentral/media_metadata/models"
	"github.com/resetcentral/media_metadata/storage"
)

func GetStudios(c *gin.Context) {
	var studios []models.Studio
	var err error

	name := c.Query("name")
	studios, err = storage.DB.FindStudios(name)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to load studios"})
		return
	}

	c.IndentedJSON(http.StatusOK, studios)
}

func GetStudioByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	studio, err := storage.DB.FindStudioByID(id)
	if err != nil || studio.ID != uint(id) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No studio found with that ID"})
		return
	}

	c.IndentedJSON(http.StatusOK, studio)
}

func PostStudio(c *gin.Context) {
	var studio models.Studio
	err := c.BindJSON(&studio)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid studio data"})
		return
	}
	if studio.Name == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "'name' field is required"})
		return
	}

	err = storage.DB.CreateStudio(&studio)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to save new studio"})
		return
	}

	c.IndentedJSON(http.StatusOK, studio)
}

func DeleteStudio(c *gin.Context) {
	var studio models.Studio

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	studio.ID = uint(id)
	err = storage.DB.DeleteStudio(studio)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete studio"})
	}
}
