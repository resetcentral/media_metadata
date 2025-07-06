package orm

import "github.com/resetcentral/media_library/models"

func (s MediaStorageGorm) CreateArtist(artists ...*models.Artist) error {
	for _, artist := range artists {
		result := s.db.Create(&artist)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (s MediaStorageGorm) FindAllArtists() ([]models.Artist, error) {
	var artists []models.Artist
	result := s.db.Find(&artists)
	return artists, result.Error
}

func (s MediaStorageGorm) FindArtistByID(id int) (models.Artist, error) {
	var artist models.Artist
	result := s.db.Find(&artist, id)
	return artist, result.Error
}

func (s MediaStorageGorm) FindArtistByName(name string) (models.Artist, error) {
	var artist models.Artist
	result := s.db.Where(&models.Artist{Name: name}).First(&artist)
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
