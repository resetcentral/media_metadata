package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/resetcentral/media_library/models"
	"github.com/resetcentral/media_library/storage"
)

func GetTags(c *gin.Context) {
	var tags []models.Tag
	var err error

	value := c.Query("value")
	tags, err = storage.DB.FindTags(value)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to load tags"})
		return
	}

	c.IndentedJSON(http.StatusOK, tags)
}

func GetTagByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	tag, err := storage.DB.FindTagByID(id)
	if err != nil || tag.ID != uint(id) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No tag found with that ID"})
		return
	}

	c.IndentedJSON(http.StatusOK, tag)
}

func PostTag(c *gin.Context) {
	var tag models.Tag
	err := c.BindJSON(&tag)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid tag data"})
		return
	}
	if tag.Value == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "'value' field is required"})
		return
	}

	err = storage.DB.CreateTag(&tag)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to save new tag"})
		return
	}

	c.IndentedJSON(http.StatusOK, tag)
}

func DeleteTag(c *gin.Context) {
	var tag models.Tag

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	tag.ID = uint(id)
	err = storage.DB.DeleteTag(tag)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete tag"})
	}
}
