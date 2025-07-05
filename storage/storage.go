package storage

import "github.com/resetcentral/media_library/models"

var DB MediaStorage

type MediaStorage interface {
	CreateArtist(...*models.Artist) error
	FindArtists() ([]models.Artist, error)
	FindArtistByID(int) (*models.Artist, error)
	DeleteArtist(...models.Artist) error

	// SaveStudio(...*Studio) error
	// FindStudios() ([]Studio, error)
	// FindStudioByID(int) (Studio, error)
	// DeleteStudio(...Studio) error

	// SaveTag(...*Tag) error
	// FindTags() ([]Tag, error)
	// FindTagByID(int) (Tag, error)
	// DeleteTag(...Tag) error

	// SaveMedia(...*AudioMedia) error
	// FindMedia() ([]AudioMedia, error)
	// FindMediaByID(int) (AudioMedia, error)
	// DeleteMedia(...AudioMedia) error
}

func NewStorage() error {
	var err error
	DB, err = NewStorageGorm()
	if err != nil {
		return err
	}

	return nil
}
