package main

import(
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
)


type Book struct{
  ID string `json:"id"`
  Title string `json:"string"`
  Author string `json:"author"`
  Quantity int  `json:"quantity"`
}

var books = []Book{
  {ID: "1", Title: "Percy Jackson and the Lighting Thief", Author:"Rick Riordan", Quantity:3},
  {ID: "2", Title: "Percy Jackson and the Sea of Monsters", Author:"Rick Riordan", Quantity:6},
  {ID: "3", Title: "Percy Jackson and the Battle of the Labyrinth", Author:"Rick Riordan", Quantity:9},
}

func getBooks(c *gin.Context) {
  c.IndentedJSON(http.StatusOK, books)
}

func addBook(c *gin.Context, b Book) {
  var nb Book
  if err:= c.BindJSON(&nb); err != nil {
    return http.StatusOK
  }
  return
}
func main() {
  router := gin.Default()
  router.GET("/books", getBooks)
  router.Run("localhost:9090")
}
