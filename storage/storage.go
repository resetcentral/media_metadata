package storage

import (
	"github.com/resetcentral/media_library/models"
	"github.com/resetcentral/media_library/storage/orm"
)

var DB MediaStorage

type MediaStorage interface {
	CreateArtist(...*models.Artist) error
	FindAllArtists() ([]models.Artist, error)
	FindArtistByID(int) (models.Artist, error)
	FindArtistByName(string) (models.Artist, error)
	DeleteArtist(...models.Artist) error

	CreateStudio(...*models.Studio) error
	FindAllStudios() ([]models.Studio, error)
	FindStudioByID(int) (models.Studio, error)
	FindStudioByName(string) (models.Studio, error)
	DeleteStudio(...models.Studio) error

	CreateTag(...*models.Tag) error
	FindAllTags() ([]models.Tag, error)
	FindTagByID(int) (models.Tag, error)
	FindTagByValue(string) (models.Tag, error)
	DeleteTag(...models.Tag) error

	CreateMedia(...*models.AudioMedia) error
	FindAllMedia() ([]models.AudioMedia, error)
	FindMediaByID(int) (models.AudioMedia, error)
	// FindMedia(models.AudioMedia) ([]models.AudioMedia, error)
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
