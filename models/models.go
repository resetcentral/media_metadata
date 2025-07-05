package models

import (
	"time"

	"gorm.io/gorm"
)

type Studio struct {
	gorm.Model
	Name string `gorm:"unique"`
}

type Artist struct {
	gorm.Model
	Name string `gorm:"unique"`
}

type Tag struct {
	gorm.Model
	Value string `gorm:"unique"`
}

type AudioMedia struct {
	gorm.Model
	Title         string
	Description   string
	FilePath      string `gorm:"unique"`
	Studio        Studio
	StudioID      int
	Artists       []Artist `gorm:"many2many:media_artists;"`
	ReleaseDate   time.Time
	Tags          []Tag `gorm:"many2many:media_tags;"`
	Length        time.Duration
	Container     string
	AudioCodec    string
	AudioChannels int
}
