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

type AudioMedia struct {
	gorm.Model
	Title         string `gorm:"index;size:256"`
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
