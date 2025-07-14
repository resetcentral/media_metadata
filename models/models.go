package models

import (
	"time"

	"gorm.io/gorm"
)

type Studio struct {
	gorm.Model
	Name string `gorm:"uniqueIndex;size:256"`
}

type Artist struct {
	gorm.Model
	Name string `gorm:"uniqueIndex;size:256"`
}

type Tag struct {
	gorm.Model
	Value string `gorm:"uniqueIndex;size:256"`
}

type AudioMetadata struct {
	gorm.Model
	AudioCodec    string
	AudioChannels uint
	MediaID       uint
}

type VideoMetadata struct {
	gorm.Model
	Width      uint
	Height     uint
	VideoCodec string
	Framerate  float64
	MediaID    uint
}

type Media struct {
	gorm.Model
	Title         string `gorm:"index;size:256"`
	Description   string
	FilePath      string `gorm:"unique"`
	Studio        Studio
	StudioID      uint
	Artists       []Artist `gorm:"many2many:media_artists;"`
	ReleaseDate   time.Time
	Tags          []Tag `gorm:"many2many:media_tags;"`
	Length        time.Duration
	Container     string
	AudioMetadata AudioMetadata
	VideoMetadata VideoMetadata
}
