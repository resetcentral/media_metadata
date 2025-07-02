package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/api/v1/media", getMedia)
	router.GET("/api/v1/media/:id", getMediaByID)
	router.POST("/api/v1/media", postMedia)
	router.PUT("/api/v1/media/:id", putMedia)
	router.PATCH("/api/v1/media/:id", patchMedia)
	router.DELETE("/api/v1/media/:id", deleteMedia)

}
