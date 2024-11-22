# Book API

This project provides a basic RESTful API for managing a collection of books. The API is built using the Go programming language and the Gin web framework. The application allows users to:

- Fetch all books
- Fetch a single book by ID
- Filter books by title
- Add new books
- Update existing books (partial updates supported)
- Checkout a book (similar to renting out)
- Checkin a book (similar to returning)

## Features

1. **GET /books**  
   Fetch all books in the collection.

2. **POST /books**  
   Add a new book to the collection.

3. **GET /books/:id**  
   Fetch a book by its ID.

4. **GET /books-filter?title={title}**  
   Filter books based on title.

5. **PATCH /books/update?id={id}**  
   Update specific fields of a book.

6. **PATCH /books/checkout?id={id}**  
   Checkout a book by reducing its quantity.

7. **PATCH /books/checkin?id={id}**  
   Checkin a book by increasing its quantity.

# Steps

### Prerequisites
- Go 1.20+ installed on your system.

### Installation
1. Clone the repository

2. To get dependencies :
    ```bash
    go mod tidy
    ```
3. To run project
    ```bash
    go run main.go
    ```
Now you can test the APIs using cURL or Postman.

HAPPY CODING!
