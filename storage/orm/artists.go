package orm

import (
	"fmt"

	"github.com/resetcentral/media_metadata/models"
)

func (s MediaStorageGorm) CreateArtist(artists ...*models.Artist) error {
	for _, artist := range artists {
		result := s.db.Create(&artist)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (s MediaStorageGorm) FindArtists(search string) ([]models.Artist, error) {
	var artists []models.Artist

	db := s.db
	if search != "" {
		search = fmt.Sprintf("%%%s%%", search)
		db = db.Where("name LIKE ?", search)
	}

	result := db.Find(&artists)
	return artists, result.Error
}

func (s MediaStorageGorm) FindArtistByID(id int) (models.Artist, error) {
	var artist models.Artist
	result := s.db.Find(&artist, id)
	return artist, result.Error
}

func (s MediaStorageGorm) DeleteArtist(artists ...models.Artist) error {
	for _, artist := range artists {
		result := s.db.Delete(&artist)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
