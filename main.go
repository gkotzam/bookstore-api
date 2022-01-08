package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Book Struck Model
type Book struct {
	ID         string   "json:\"id\""
	Title      string   "json:\"title\""
	Isbn10     string   "json:\"isbn10\""
	Isbn13     string   "json:\"isbn13\""
	Authors    []Author "json:\"authors\""
	Publisher  string   "json:\"publisher\""
	Pub_year   string   "json:\"pub_year\""
	Language   string   "json:\"language\""
	Pages      string   "json:\"pages\""
	Decription string   "json:\"description\""
	Price      string   "json:\"price\""
}

// Author Struck Model
type Author struct {
	ID      string "json:\"author_id\""
	Name    string "json:\"name\""
	Country string "json:\"country\""
	//Written_books string "json:\"written_books\""
}

// BOOKS

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	books := db_get_books()
	json.NewEncoder(w).Encode(books)
}

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	num, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err.Error())
	}
	book := db_get_book(num)
	json.NewEncoder(w).Encode(book)
}

// Create a new Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	db_create_book(book)
}

// Update a Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	num, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err.Error())
	}
	db_update_book(num, book)
}

// Delete a Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	num, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err.Error())
	}
	db_delete_book(num)
}

// AUTHORS

// Get All Authors
func getAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authors := db_get_authors()
	json.NewEncoder(w).Encode(authors)
}

// Get Single Author
func getAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	author := db_get_author_by_id(params["id"])
	json.NewEncoder(w).Encode(author)
}

// Create a new Author
func createAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var author Author
	_ = json.NewDecoder(r.Body).Decode(&author)
	db_create_author(author)
}

// Update an Author
func updateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var author Author
	_ = json.NewDecoder(r.Body).Decode(&author)
	db_update_author(params["id"], author)
}

// Delete an Author
func deleteAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	num, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err.Error())
	}
	db_delete_author(num)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "homepage Endpoint Hit")
}

func main() {

	// Init Router
	router := mux.NewRouter()

	// Route Handlers / Endpoint
	router.HandleFunc("/", homePage)

	// BOOKS
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PATCH")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	// AUTHORS
	router.HandleFunc("/api/authors", getAuthors).Methods("GET")
	router.HandleFunc("/api/authors/{id}", getAuthor).Methods("GET")
	router.HandleFunc("/api/authors", createAuthor).Methods("POST")
	router.HandleFunc("/api/authors/{id}", updateAuthor).Methods("PATCH")
	router.HandleFunc("/api/authors/{id}", deleteAuthor).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}
