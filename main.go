package main

import (
	"my-rest-api/repositories"
	"my-rest-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAlbums(albumRepo repositories.IAlbumRepository) func(*gin.Context) {
	return func(c *gin.Context) {
		var albums []repositories.Album = albumRepo.GetAllAlbums()
		c.IndentedJSON(http.StatusOK, albums)
	}
}

func ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, utils.NewSuccessResponse(
		"Lets go 🚀🚀",
		nil),
	)
}

func main() {
	albumRepo := &repositories.AlbumRepository{}
	router := gin.Default()
	router.GET("/", ping)
	router.GET("/albums", getAlbums(albumRepo))
	router.Run("localhost:3000")
}
