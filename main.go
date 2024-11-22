package main

import (
	"book-api/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//basic router
	router := gin.Default()
	fmt.Println("Running server./././.")

	//routes
	router.GET("/books", controllers.GetBooks)
	router.POST("/books", controllers.CreateBook)
	router.GET("/books-filter", controllers.FilterBookByTitle)
	router.GET("/books/:id", controllers.GetBookById)
	router.PATCH("/books/checkout", controllers.CheckOutBooks)
	router.PATCH("/books/checkin", controllers.CheckInBooks)
	router.PATCH("/books/update", controllers.UpdateBook)
	//running server
	router.Run("localhost:8080")
}
