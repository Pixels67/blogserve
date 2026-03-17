package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var db Database

func main() {
	err := db.Init()
	if err != nil {
		panic(err)
	}

	defer db.Deinit()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.GET("/posts/:file", retrieve)
	router.GET("/posts", retrieveAll)
	router.Run("0.0.0.0:" + port)
}

func retrieve(c *gin.Context) {
	file := c.Param("file")
	record, err := db.Get(file)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found"})
		return
	}

	post, err := RecordToPost(file, record)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, post)
}

func retrieveAll(c *gin.Context) {
	records, err := db.GetAll()

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found"})
		return
	}

	posts := []Post{}
	for file, record := range records {
		post, err := RecordToPost(file, record)
		if err != nil {
			continue
		}

		posts = append(posts, post)
	}

	c.IndentedJSON(http.StatusOK, posts)
}
