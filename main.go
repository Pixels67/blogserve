package main

import (
	"fmt"
	"net/http"
	"os"

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.GET("/posts/:file", retrieve)
	router.Run("0.0.0.0:" + port)
}

func retrieve(c *gin.Context) {
	file := c.Param("file")
	record, err := db.Get(file)

	fmt.Println(record)

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
