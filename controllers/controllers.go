package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// creating a model for data
type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

// random data
var books = []book{
	{ID: "1", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 3},
	{ID: "2", Title: "To Kill a Mockingbird", Author: "Harper Lee", Quantity: 5},
	{ID: "3", Title: "1984", Author: "George Orwell", Quantity: 7},
}

// controller to GET books
func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// controller to CREATE books
func CreateBook(c *gin.Context) {
	var newBook book

	err := c.BindJSON(&newBook)
	if err != nil {
		fmt.Println("Error occured in binding")
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// utility function to get book by id
func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

// controller to get book by ID
func GetBookById(c *gin.Context) {
	var id string = c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

// controller to find a book by title
func FilterBookByTitle(c *gin.Context) {
	title, ok := c.GetQuery("title")
	var resultBook []book
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid title"})
		return
	}

	for i, b := range books {
		if strings.Contains(strings.ToLower(b.Title), strings.ToLower(title)) {
			resultBook = append(resultBook, books[i])
		}
	}

	c.IndentedJSON(http.StatusOK, resultBook)
}

// controller to update book
func UpdateBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	var existingBook *book
	for i := range books {
		if books[i].ID == id {
			existingBook = &books[i]
			break
		}
	}

	var newBookData map[string]interface{}
	if err := c.BindJSON(&newBookData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse the data"})
		return
	}

	// Update only the provided fields
	if title, ok := newBookData["title"].(string); ok {
		existingBook.Title = title
	}
	if author, ok := newBookData["author"].(string); ok {
		existingBook.Author = author
	}
	if quantity, ok := newBookData["quantity"].(float64); ok { // JSON numbers are float64 by default
		existingBook.Quantity = int(quantity)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Updated!"})
}

// controller to checkout book
func CheckOutBooks(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if ok == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "incorrect id"})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "book not available"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

// controller to checkin book
func CheckInBooks(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if ok == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}
