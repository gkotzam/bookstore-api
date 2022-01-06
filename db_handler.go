package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// db connection
func db_conn() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/api-test")

	if err != nil {
		log.Panic(err.Error())
	}
	return db
}

// Get all books
func db_get_books() (books []Book) {

	db := db_conn()
	defer db.Close()

	var sql_query = "SELECT * from books;"

	result, err := db.Query(sql_query)

	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	var book Book

	for result.Next() {
		err = result.Scan(&book.ID, &book.Title, &book.Isbn10, &book.Isbn13, &book.Author, &book.Publisher, &book.Pub_year, &book.Language, &book.Pages, &book.Decription, &book.Price)
		if err != nil {
			panic(err.Error())
		}
		books = append(books, book)
	}

	return books

}

// Get single book
func db_get_book(num int) (book Book) {
	db := db_conn()
	defer db.Close()

	sql_query := fmt.Sprintf("SELECT * from books WHERE id=%d;", num)

	result, err := db.Query(sql_query)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	if result.Next() {
		err = result.Scan(&book.ID, &book.Title, &book.Isbn10, &book.Isbn13, &book.Author, &book.Publisher, &book.Pub_year, &book.Language, &book.Pages, &book.Decription, &book.Price)
		if err != nil {
			panic(err.Error())
		}
	}

	return book

}
