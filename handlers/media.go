package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/resetcentral/media_metadata/models"
	"github.com/resetcentral/media_metadata/storage"
)

func PostMedia(c *gin.Context) {
	var media models.Media
	err := c.BindJSON(&media)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid media data"})
		return
	}

	err = storage.DB.CreateMedia(&media)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to save new media"})
		return
	}

	c.IndentedJSON(http.StatusOK, media)
}
