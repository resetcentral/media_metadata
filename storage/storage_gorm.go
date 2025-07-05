package storage

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/resetcentral/media_library/models"
)

type MediaStorageGorm struct {
	db *gorm.DB
}
type DBConfig struct {
	Host     string
	Name     string
	User     string
	Password string
}

func NewStorageGorm() (MediaStorageGorm, error) {
	var storage MediaStorageGorm
	var dbConfig DBConfig

	err := envconfig.Process("db", &dbConfig)
	if err != nil {
		return storage, err
	}
	dsnFmt := "%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(dsnFmt, dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return storage, err
	}

	err = db.AutoMigrate(&models.Artist{})
	// err = db.AutoMigrate(&AudioMedia{})

	if err != nil {
		return storage, err
	}

	storage.db = db
	return storage, nil
}

func (s MediaStorageGorm) CreateArtist(artists ...*models.Artist) error {
	for _, artist := range artists {
		result := s.db.Create(artist)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
func (s MediaStorageGorm) FindArtists() ([]models.Artist, error) {
	var artists []models.Artist
	result := s.db.Find(&artists)
	if result.Error != nil {
		return nil, result.Error
	}

	return artists, nil
}
func (s MediaStorageGorm) FindArtistByID(id int) (*models.Artist, error) {
	var artist models.Artist
	result := s.db.Find(&artist, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &artist, nil

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

func (s MediaStorageGorm) CreateStudio(studios ...*models.Studio) error {
	for _, studio := range studios {
		result := s.db.Create(studio)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (s MediaStorageGorm) FindStudios() ([]models.Studio, error) {
	var studios []models.Studio
	result := s.db.Find(&studios)
	if result.Error != nil {
		return nil, result.Error
	}

	return studios, nil
}

func (s MediaStorageGorm) FindStudioByID(id int) (*models.Studio, error) {
	var studio models.Studio
	result := s.db.Find(&studio, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &studio, nil
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

// func (s *StorageGorm) SaveTag(...*Tag) error {

// }
// func (s *StorageGorm) FindTags() ([]Tag, error) {

// }
// func (s *StorageGorm) FindTagByID(int) (Tag, error) {

// }
// func (s *StorageGorm) DeleteTag(...Tag) error {

// }

// func (s *StorageGorm) SaveMedia(...*AudioMedia) error {

// }
// func (s *StorageGorm) FindMedia() ([]AudioMedia, error) {

// }
// func (s *StorageGorm) FindMediaByID(int) (AudioMedia, error) {

// }
// func (s *StorageGorm) DeleteMedia(...AudioMedia) error {

// }
