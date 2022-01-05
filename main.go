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
	ID     string "json:\"id\""
	Isbn   string "json:\"isbn\""
	Title  string "json:\"title\""
	Author string "json:\"author\""
}

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

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "homepage Endpoint Hit")
}

func main() {

	// Init Router
	router := mux.NewRouter()

	// Route Handlers / Endpoint
	router.HandleFunc("/", homePage)
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))

}
