package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Artist struct {
	gorm.Model
	Name string
}

type Tag struct {
	gorm.Model
	Value string
}

type AudioMedia struct {
	gorm.Model
	Title         string
	Description   string
	FilePath      string
	Studio        string
	Artists       []Artist `gorm:"many2many:media_artists;"`
	ReleaseDate   time.Time
	Tags          []Tag `gorm:"many2many:media_tags;"`
	Length        time.Duration
	Container     string
	AudioCodec    string
	AudioChannels int
	MediaType     string
	MediaID       int
}

type VideoMedia struct {
	gorm.Model
	Audio      AudioMedia `gorm:"polymorphic:Media;"`
	Width      int
	Height     int
	VideoCodec string
	Fps        int
}

func getMedia(c *gin.Context) {
}

func getMediaByID(c *gin.Context) {
	// id := c.Param("id")
}

func postMedia(c *gin.Context) {

}

func putMedia(c *gin.Context) {
	// id := c.Param("id")
}

func patchMedia(c *gin.Context) {
	// id := c.Param("id")
}

func deleteMedia(c *gin.Context) {
	// id := c.Param("id")
}
