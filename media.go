package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AudioMedia struct {
	gorm.Model
	FilePath      string
	Length        time.Duration
	Studio        string
	Artists       []string
	Tags          []string
	Title         string
	Description   string
	Container     string
	AudioCodec    string
	AudioChannels int
}

type VideoMedia struct {
	AudioMedia
	Dimensions [2]int
	VideoCodec string
	Fps        int
}

func getMedia(c *gin.Context) {

}

func getMediaByID(c *gin.Context) {
	id := c.Param("id")
}

func postMedia(c *gin.Context) {

}

func putMedia(c *gin.Context) {
	id := c.Param("id")
}

func patchMedia(c *gin.Context) {
	id := c.Param("id")
}

func deleteMedia(c *gin.Context) {
	id := c.Param("id")
}
