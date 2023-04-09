package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"rest-go/models"

	"github.com/gin-gonic/gin"
)

var albums = []models.Album{
	{ID: "1", Title: "RAIZ", Artist: "João Gomes", Price: 56.99},
	{ID: "2", Title: "LITTLE LOVE", Artist: "MC Cabelinho", Price: 17.99},
	{ID: "3", Title: "Modo Repeat", Artist: "Felipe Amorim", Price: 39.99},
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumsByID(c *gin.Context) {

	id := c.Param("id")

	for _, x := range albums {
		if x.ID == id {
			c.IndentedJSON(http.StatusOK, x)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func PostAlbums(c *gin.Context) {

	var newAlbum models.Album

	reqBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading request body: %v", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = json.Unmarshal(reqBody, &newAlbum)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing request body: %v", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func PutAlbums(c *gin.Context) {

	id := c.Param("id")

	var updatedAlbum models.Album

	reqBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading request body: %v", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = json.Unmarshal(reqBody, &updatedAlbum)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing request body: %v", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, x := range albums {
		if x.ID == id {
			albums[i] = updatedAlbum
			c.IndentedJSON(http.StatusOK, updatedAlbum)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func DeleteAlbums(c *gin.Context) {
	
	id := c.Param("id")

	for i, x := range albums {
		if x.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
