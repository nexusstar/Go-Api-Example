package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "gopkg.in/validator.v2"
  "errors"
  )

type book struct {
  ID string `json:"id"`
  Title string `json:"title" validate:"nonzero"`
  Author string `json:"author" validate:"nonzero"`
  Description string `json:"description" validate:"nonzero"`
  Quantity int `json:"quantity" validate:"nonzero"`
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

func getBookById(id string) (*book, error) {
  for i, book := range books {
    if book.ID == id {
      return &books[i], nil
    }
  }
  return nil, errors.New("Book not found")
}

func getBook(c *gin.Context) {
  id := c.Param("id")
  book, err := getBookById(id)
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "message": "Book not found",
    })
    return
  }
  c.IndentedJSON(http.StatusOK, book)
}

func addBook(c *gin.Context) {
  var newBook book
  
  if err := c.BindJSON(&newBook); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "message": "Invalid request",
    })
    return
  }

  if err:= validator.Validate(newBook); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }

  books = append(books, newBook)
  c.JSON(http.StatusCreated, gin.H{
    "message": "Book added successfully",
    "book": newBook,
  })
}

func updateBook(c *gin.Context) {
  id := c.Param("id")

  var newBook book
  
  err := c.BindJSON(&newBook); if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "message": "Invalid request",
    })
    return
  }

  if err:= validator.Validate(newBook); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }

  book, err := getBookById(id); if err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "message": "Book not found",
    })
    return
  }

  if(newBook.Title != book.Title && newBook.Title != ""){
    book.Title = newBook.Title
  }

  if(newBook.Author != book.Author && newBook.Author != ""){
    book.Author = newBook.Author
  }

  if(newBook.Description != book.Description && newBook.Description != ""){
    book.Description = newBook.Description
  }

  if(newBook.Quantity != book.Quantity){
    book.Quantity = newBook.Quantity
  }

  c.JSON(http.StatusOK, gin.H{
    "message": "Book updated successfully",
    "book": &book,
  })
}

func main () {
  router := gin.Default()
  router.GET("/", getApi)
  router.GET("/books", getBooks)
  router.GET("/books/:id", getBook)
  router.POST("/books", addBook)
  router.PUT("/books/:id", updateBook)
  // router.DELETE("/books/:id", deleteBook)
  router.Run(":8080")
}
