package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Name     string
	User     string
	Password string
}

func main() {
	var dbConfig DBConfig
	err := envconfig.Process("db", &dbConfig)
	if err != nil {
		log.Fatal(err.Error())
	}
	dsnFmt := "%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(dsnFmt, dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to DB %v\n", err)
	}

	db.AutoMigrate(&AudioMedia{})
	db.AutoMigrate(&VideoMedia{})

	router := gin.Default()
	router.GET("/api/v1/media", getMedia)
	router.GET("/api/v1/media/:id", getMediaByID)
	router.POST("/api/v1/media", postMedia)
	router.PUT("/api/v1/media/:id", putMedia)
	router.PATCH("/api/v1/media/:id", patchMedia)
	router.DELETE("/api/v1/media/:id", deleteMedia)

}
