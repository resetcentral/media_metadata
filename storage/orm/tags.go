package orm

import "github.com/resetcentral/media_library/models"

func (s MediaStorageGorm) CreateTag(tags ...*models.Tag) error {
	for _, tag := range tags {
		result := s.db.Create(tag)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (s MediaStorageGorm) FindAllTags() ([]models.Tag, error) {
	var tags []models.Tag
	result := s.db.Find(&tags)
	return tags, result.Error
}

func (s MediaStorageGorm) FindTagByID(id int) (models.Tag, error) {
	var tag models.Tag
	result := s.db.Find(&tag, id)
	return tag, result.Error
}

func (s MediaStorageGorm) FindTagByValue(value string) (models.Tag, error) {
	var tag models.Tag
	result := s.db.Where(&models.Tag{Value: value}).First(&tag)
	return tag, result.Error
}

func (s MediaStorageGorm) DeleteTag(tags ...models.Tag) error {
	for _, tag := range tags {
		result := s.db.Delete(&tag)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
