package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  // "errors"
  )

type book struct {
  ID string `json:"id"`
  Title string `json:"title"`
  Author string `json:"author"`
  Description string `json:"description"`
  Quantity int `json:"quantity"`
}

var books = []book{
  {ID: "1", Title: "Golang pointers", Author: "John Doe", Description: "Golang pointers", Quantity: 10},
  {ID: "2", Title: "Golang arrays", Author: "Jane Doe", Description: "Golang arrays", Quantity: 4},
  {ID: "3", Title: "Golang slices", Author: "Janet Doe", Description: "Golang slices", Quantity: 20},
}

func getApi(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
    "message": "Hello, this is example Go API",
  })
}

func getBooks(c *gin.Context) {
  c.IndentedJSON(http.StatusOK, books)
}

func main () {
  router := gin.Default()
  router.GET("/", getApi)
  router.GET("/books", getBooks)
  // router.GET("/books/:id", getBook)
  // router.POST("/books", addBook)
  // router.PUT("/books/:id", updateBook)
  // router.DELETE("/books/:id", deleteBook)
  router.Run(":8080")
}
