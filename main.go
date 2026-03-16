package main

import (
	"fmt"
	"net/http"

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

	router.GET("/posts/:file", retrieve)
	router.Run("localhost:8080")
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
