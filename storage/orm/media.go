package orm

import (
	"github.com/resetcentral/media_library/models"
)

func (s MediaStorageGorm) CreateMedia(medias ...*models.AudioMedia) error {
	for _, media := range medias {
		result := s.db.Create(&media)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (s MediaStorageGorm) FindMedia() ([]models.AudioMedia, error) {
	var media []models.AudioMedia
	result := s.db.Find(&media)
	return media, result.Error
}

func (s MediaStorageGorm) FindMediaByID(id int) (models.AudioMedia, error) {
	var media models.AudioMedia
	result := s.db.Find(&media, id)
	return media, result.Error
}

func (s MediaStorageGorm) DeleteMedia(medias ...models.AudioMedia) error {
	for _, media := range medias {
		result := s.db.Delete(&media)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
