package router

import (
	"rest-go/controller"

	"github.com/gin-gonic/gin"
)

func StartRouter() {
	router := gin.Default()
	router.GET("/albums", controller.GetAlbums)
	router.GET("/albums/:id", controller.GetAlbumsByID)
	router.POST("/albums", controller.PostAlbums)
	router.PUT("/albums/:id", controller.PutAlbums)
	router.DELETE("/albums/:id", controller.DeleteAlbums)
	router.Run("localhost:8080")
}