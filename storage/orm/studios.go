package orm

import (
	"fmt"

	"github.com/resetcentral/media_metadata/models"
)

func (s MediaStorageGorm) CreateStudio(studios ...*models.Studio) error {
	for _, studio := range studios {
		result := s.db.Create(studio)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (s MediaStorageGorm) FindStudios(search string) ([]models.Studio, error) {
	var studios []models.Studio
	db := s.db
	if search != "" {
		search = fmt.Sprintf("%%%s%%", search)
		db = db.Where("name LIKE ?", search)
	}

	result := db.Find(&studios)
	return studios, result.Error
}

func (s MediaStorageGorm) FindStudioByID(id int) (models.Studio, error) {
	var studio models.Studio
	result := s.db.Find(&studio, id)
	return studio, result.Error
}

func (s MediaStorageGorm) DeleteStudio(studios ...models.Studio) error {
	for _, studio := range studios {
		result := s.db.Delete(&studio)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
