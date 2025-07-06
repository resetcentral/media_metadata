package orm

import "github.com/resetcentral/media_library/models"

func (s MediaStorageGorm) CreateStudio(studios ...*models.Studio) error {
	for _, studio := range studios {
		result := s.db.Create(studio)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (s MediaStorageGorm) FindAllStudios() ([]models.Studio, error) {
	var studios []models.Studio
	result := s.db.Find(&studios)
	return studios, result.Error
}

func (s MediaStorageGorm) FindStudioByID(id int) (models.Studio, error) {
	var studio models.Studio
	result := s.db.Find(&studio, id)
	return studio, result.Error
}

func (s MediaStorageGorm) FindStudioByName(name string) (models.Studio, error) {
	var studio models.Studio
	result := s.db.Where(&models.Studio{Name: name}).First(&studio)
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
