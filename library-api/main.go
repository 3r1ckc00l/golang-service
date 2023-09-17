package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)
	router.GET("/books/:id/:title", booksHandler)
	router.GET("/query", queryHandler)
	router.POST("/books", postBooksHandler)

	router.Run(":5000")
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title": "golang GIN",
		"desc":  "using Gonic framework",
	})
}
func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title": "Hello World",
		"desc":  "lorem ipsum dolor sit amet",
	})
}
func booksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}
func queryHandler(c *gin.Context) {
	get := c.Query("v")

	c.JSON(http.StatusOK, gin.H{
		"data": get,
	})
}

type BookInput struct {
	Title string
	Price int
}

func postBooksHandler(c *gin.Context) {
	var bookInput BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
	})
}
