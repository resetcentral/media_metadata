package storage

import (
	"github.com/resetcentral/media_metadata/models"
	"github.com/resetcentral/media_metadata/storage/orm"
)

var DB MediaStorage

type MediaStorage interface {
	CreateArtist(...*models.Artist) error
	FindArtists(string) ([]models.Artist, error)
	FindArtistByID(int) (models.Artist, error)
	DeleteArtist(...models.Artist) error

	CreateStudio(...*models.Studio) error
	FindStudios(string) ([]models.Studio, error)
	FindStudioByID(int) (models.Studio, error)
	DeleteStudio(...models.Studio) error

	CreateTag(...*models.Tag) error
	FindTags(string) ([]models.Tag, error)
	FindTagByID(int) (models.Tag, error)
	DeleteTag(...models.Tag) error

	CreateMedia(...*models.AudioMedia) error
	FindMedia() ([]models.AudioMedia, error)
	FindMediaByID(int) (models.AudioMedia, error)
	DeleteMedia(...models.AudioMedia) error
}

func NewStorage() error {
	var err error
	DB, err = orm.NewStorageGorm()
	if err != nil {
		return err
	}

	return nil
}
