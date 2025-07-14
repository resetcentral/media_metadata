package orm

import (
	"github.com/resetcentral/media_metadata/models"
)

func (s MediaStorageGorm) CreateMedia(medias ...*models.Media) error {
	for _, media := range medias {
		result := s.db.Create(&media)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (s MediaStorageGorm) FindMedia() ([]models.Media, error) {
	var media []models.Media
	result := s.db.Preload("AudioMetadata").Preload("VideoMetadata").Find(&media)
	return media, result.Error
}

func (s MediaStorageGorm) FindMediaByID(id int) (models.Media, error) {
	var media models.Media
	result := s.db.Preload("AudioMetadata").Preload("VideoMetadata").Find(&media, id)
	return media, result.Error
}

func (s MediaStorageGorm) DeleteMedia(medias ...models.Media) error {
	for _, media := range medias {
		result := s.db.Delete(&media)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
