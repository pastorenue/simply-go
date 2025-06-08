package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

// Book represents book data
type Book struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Author string  `json:"author"`
    Price  float64 `json:"price"`
}

// GetBooks returns all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    // Sample data
    books := []Book{
        {
            ID:     "1",
            Title:  "The Go Programming Language",
            Author: "Alan Donovan",
            Price:  29.99,
        },
        {
            ID:     "2",
            Title:  "Go Web Programming",
            Author: "Sau Sheong Chang",
            Price:  24.99,
        },
    }
    
    json.NewEncoder(w).Encode(books)
}

// GetBook returns a single book
func GetBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    // Get URL parameters
    params := mux.Vars(r)
    
    // Sample data - in real app, would fetch from database
    book := Book{
        ID:     params["id"],
        Title:  "The Go Programming Language",
        Author: "Alan Donovan",
        Price:  29.99,
    }
    
    json.NewEncoder(w).Encode(book)
}

func Handler() {
    // Initialize router
    r := mux.NewRouter()
    
    // Route handlers
    r.HandleFunc("/api/books", GetBooks).Methods("GET")
    r.HandleFunc("/api/books/{id}", GetBook).Methods("GET")
    
    // Start server
    log.Fatal(http.ListenAndServe(":1212", r))
}