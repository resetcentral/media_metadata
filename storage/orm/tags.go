package orm

import (
	"fmt"

	"github.com/resetcentral/media_metadata/models"
)

func (s MediaStorageGorm) CreateTag(tags ...*models.Tag) error {
	for _, tag := range tags {
		result := s.db.Create(tag)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (s MediaStorageGorm) FindTags(search string) ([]models.Tag, error) {
	var tags []models.Tag
	db := s.db
	if search != "" {
		search = fmt.Sprintf("%%%s%%", search)
		db = db.Where("value LIKE ?", search)
	}

	result := db.Find(&tags)
	return tags, result.Error
}

func (s MediaStorageGorm) FindTagByID(id int) (models.Tag, error) {
	var tag models.Tag
	result := s.db.Find(&tag, id)
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
