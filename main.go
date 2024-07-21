package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

//album slice to seed initial record album data
var albums = []album{
	{ID: "1", Title:"Blue Train", Artist: "John Coltrane", Price: 64.99},
	{ID: "2", Title:"Red Train", Artist: "John Boltrane", Price: 64.99},
	{ID: "3", Title:"Tellow Train", Artist: "John Soltrane", Price: 64.99},
	{ID: "4", Title:"Green Train", Artist: "John Woltrane", Price: 64.99},
	{ID: "5", Title:"Orange Train", Artist: "John Koltrane", Price: 64.99},
}

//get albums -> get all albums
func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums)
}

//post albums -> post request will add
func postAlbums(c *gin.Context){
	var newAlbum album
	//call bind json to bind received json to new album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	//add new albums
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

//get album by id 
func getAlbumByID(c *gin.Context){
	id := c.Param("id")

	//loop over list of albums looking for album whose id value matches parameter
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"album not found!"})

}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}