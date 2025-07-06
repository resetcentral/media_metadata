package orm

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

	db.AutoMigrate(&models.Artist{})
	db.AutoMigrate(&models.Studio{})
	db.AutoMigrate(&models.Tag{})
	err = db.AutoMigrate(&models.AudioMedia{})

	if err != nil {
		return storage, err
	}

	storage.db = db
	return storage, nil
}

// func (s *StorageGorm) SaveMedia(...*AudioMedia) error {

// }
// func (s *StorageGorm) FindMedia() ([]AudioMedia, error) {

// }
// func (s *StorageGorm) FindMediaByID(int) (AudioMedia, error) {

// }
// func (s *StorageGorm) DeleteMedia(...AudioMedia) error {

// }
