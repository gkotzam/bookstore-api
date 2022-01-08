package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

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

// Adds an author with (@param: name) in database if author does not exists
func db_add_author(name string) {

	if db_get_author_id(name) == -1 {
		db := db_conn()
		defer db.Close()

		sql_query := fmt.Sprintf("INSERT INTO authors (name) VALUES ('%s');", name)

		result, err := db.Query(sql_query)

		if err != nil {
			panic(err.Error())
		}
		defer result.Close()
	}

}

// Returns author's id with (@param: name)
//
// IF author does not exists return -1
func db_get_author_id(name string) (id int) {
	db := db_conn()
	defer db.Close()

	sql_query := fmt.Sprintf("SELECT id from authors WHERE name='%s';", name)

	result, err := db.Query(sql_query)

	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	if result.Next() {
		err = result.Scan(&id)
		if err != nil {
			panic(err.Error())
		}
	} else {
		id = -1
	}
	return id
}

// Returns Author with id = (@param id)
func db_get_author_by_id(id string) (author Author) {
	db := db_conn()
	defer db.Close()

	sql_query := fmt.Sprintf("SELECT * from authors WHERE id='%s';", id)

	result, err := db.Query(sql_query)

	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	if result.Next() {
		err = result.Scan(&author.ID, &author.Name, &author.Country)
		if err != nil {
			panic(err.Error())
		}
	}
	return author
}

// Rerurns all books stored in database
func db_get_books() (books []Book) {

	db := db_conn()
	defer db.Close()

	var sql_query = "SELECT * from books;"

	result, err := db.Query(sql_query)

	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	var temp string

	for result.Next() {
		var book Book
		err = result.Scan(&book.ID, &book.Title, &book.Isbn10, &book.Isbn13, &temp, &book.Publisher, &book.Pub_year, &book.Language, &book.Pages, &book.Decription, &book.Price)
		if err != nil {
			panic(err.Error())
		}

		temp_authors := strings.Split(temp, ",")

		for _, author_id := range temp_authors {
			book.Authors = append(book.Authors, db_get_author_by_id(author_id))
		}

		books = append(books, Book{ID: book.ID, Title: book.Title, Isbn10: book.Isbn10, Isbn13: book.Isbn13, Authors: book.Authors, Publisher: book.Publisher, Pub_year: book.Pub_year, Language: book.Language, Pages: book.Pages, Decription: book.Decription, Price: book.Price})

	}

	return books

}

// Returns a single book with ID = @param: num
//
// Example: book := db_get_book(1) , return the book with id = 1
func db_get_book(num int) (book Book) {
	db := db_conn()
	defer db.Close()

	sql_query := fmt.Sprintf("SELECT * from books WHERE id=%d;", num)

	result, err := db.Query(sql_query)

	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	var temp string

	if result.Next() {
		err = result.Scan(&book.ID, &book.Title, &book.Isbn10, &book.Isbn13, &temp, &book.Publisher, &book.Pub_year, &book.Language, &book.Pages, &book.Decription, &book.Price)
		if err != nil {
			panic(err.Error())
		}

		temp_authors := strings.Split(temp, ",")

		for _, author_id := range temp_authors {
			book.Authors = append(book.Authors, db_get_author_by_id(author_id))
		}

		//book = Book{ID: book.ID, Title: book.Title, Isbn10: book.Isbn10, Isbn13: book.Isbn13, Authors: book.Authors, Publisher: book.Publisher, Pub_year: book.Pub_year, Language: book.Language, Pages: book.Pages, Decription: book.Decription, Price: book.Price}
	}

	return book

}

// Returns a string which contains all authors ids sep ,
func db_authors_to_id_string(authors []Author) (id_string string) {
	db := db_conn()
	defer db.Close()

	id_string = ""
	// get all ids in one string sep ,
	for _, author := range authors {
		id_string += fmt.Sprintf("%d,", db_get_author_id(author.Name))
	}
	if len(id_string) > 0 {
		id_string = id_string[:len(id_string)-1] //remove last ,
	}
	return id_string

}

// Stores a book in database
func db_create_book(book Book) {
	db := db_conn()
	defer db.Close()
	// add slashes
	book.Title = strings.ReplaceAll(book.Title, "'", "\\'")
	for _, author := range book.Authors {
		author.Name = strings.ReplaceAll(author.Name, "'", "\\'")
		// adds author in database
		db_add_author(author.Name)
	}
	book.Publisher = strings.ReplaceAll(book.Publisher, "'", "\\'")
	book.Language = strings.ReplaceAll(book.Language, "'", "\\'")
	book.Decription = strings.ReplaceAll(book.Decription, "'", "\\'")

	author_ids := db_authors_to_id_string(book.Authors)

	sql_query := fmt.Sprintf("INSERT INTO books ( title, isbn10, isbn13, author, publisher, pub_year, language, pages, description, price) VALUES ('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s');", book.Title, book.Isbn10, book.Isbn13, author_ids, book.Publisher, book.Pub_year, book.Language, book.Pages, book.Decription, book.Price)

	result, err := db.Query(sql_query)

	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
}

// Updates a book in database
func db_update_book(num int, book Book) {
	db := db_conn()
	defer db.Close()

	old_book := db_get_book(num)

	var author_ids string

	if len(book.Authors) == 0 {
		// keep old book info
		author_ids = db_authors_to_id_string(old_book.Authors)
	} else { // update
		for _, author := range book.Authors {
			author.Name = strings.ReplaceAll(author.Name, "'", "\\'")
			// adds author in database
			db_add_author(author.Name)
		}
		author_ids = db_authors_to_id_string(book.Authors)
	}

	// Keep old book info or update

	if book.Title == "" {
		book.Title = old_book.Title
	} else { // add slashes
		book.Title = strings.ReplaceAll(book.Title, "'", "\\'")
	}
	if book.Isbn10 == "" {
		book.Isbn10 = old_book.Isbn10
	}
	if book.Isbn13 == "" {
		book.Isbn13 = old_book.Isbn13
	}
	if book.Publisher == "" {
		book.Publisher = old_book.Publisher
	} else { // add slashes
		book.Publisher = strings.ReplaceAll(book.Publisher, "'", "\\'")
	}
	if book.Pub_year == "" {
		book.Pub_year = old_book.Pub_year
	}
	if book.Language == "" {
		book.Language = old_book.Language
	} else { // add slashes
		book.Language = strings.ReplaceAll(book.Language, "'", "\\'")
	}
	if book.Pages == "" {
		book.Pages = old_book.Pages
	}
	if book.Decription == "" {
		book.Decription = old_book.Decription
	} else { // add slashes
		book.Decription = strings.ReplaceAll(book.Decription, "'", "\\'")
	}
	if book.Price == "" {
		book.Price = old_book.Price
	}

	sql_query := fmt.Sprintf("UPDATE books SET title ='%s',  isbn10 ='%s',  isbn13 ='%s',  author ='%s',  publisher ='%s',  pub_year ='%s',  language ='%s',  pages ='%s',  description ='%s',  price ='%s' WHERE id=%d;", book.Title, book.Isbn10, book.Isbn13, author_ids, book.Publisher, book.Pub_year, book.Language, book.Pages, book.Decription, book.Price, num)

	result, err := db.Query(sql_query)

	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
}

// Deletes a book with ID = num
//
// Example: db_delete_book(1) , deletes the book with id = 1
func db_delete_book(num int) {
	db := db_conn()
	defer db.Close()

	sql_query := fmt.Sprintf("DELETE from books WHERE id=%d;", num)

	result, err := db.Query(sql_query)

	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	// get all books after the deleted book
	sql_query = fmt.Sprintf("SELECT * from books WHERE id >%d;", num)

	result2, err := db.Query(sql_query)

	if err != nil {
		panic(err.Error())
	}

	defer result2.Close()

	var book Book
	var temp string
	for result2.Next() {
		// reset id for every book after the deleted book
		err = result2.Scan(&book.ID, &book.Title, &book.Isbn10, &book.Isbn13, &temp, &book.Publisher, &book.Pub_year, &book.Language, &book.Pages, &book.Decription, &book.Price)
		if err != nil {
			panic(err.Error())
		}
		id, err := strconv.Atoi(book.ID)
		if err != nil {
			panic(err.Error())
		}
		var newId int = id - 1
		// set id-1
		sql_query = fmt.Sprintf("UPDATE books SET id = %d WHERE id=%d;", newId, id)
		result3, err := db.Query(sql_query)
		if err != nil {
			panic(err.Error())
		}
		result3.Close()

	}

	// reset auto_increment
	sql_query = "SELECT COUNT(1) from books;"
	result4, err := db.Query(sql_query)
	if err != nil {
		panic(err.Error())
	}
	defer result4.Close()
	if result4.Next() {
		var total int
		err = result4.Scan(&total)
		if err != nil {
			panic(err.Error())
		}
		sql_query = fmt.Sprintf("ALTER TABLE books AUTO_INCREMENT = %d;", total)
		result5, err := db.Query(sql_query)
		if err != nil {
			panic(err.Error())
		}
		defer result5.Close()
	}

}

// Returns all authors stored in database
func db_get_authors() (authors []Author) {
	db := db_conn()
	defer db.Close()

	var sql_query = "SELECT * from authors;"

	result, err := db.Query(sql_query)

	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var author Author
		err = result.Scan(&author.ID, &author.Name, &author.Country)
		if err != nil {
			panic(err.Error())
		}

		authors = append(authors, author)

	}

	return authors
}

// Stores an author in database
func db_create_author(author Author) {
	db := db_conn()
	defer db.Close()
	// add slashes
	author.Country = strings.ReplaceAll(author.Country, "'", "\\'")
	author.Name = strings.ReplaceAll(author.Name, "'", "\\'")

	sql_query := fmt.Sprintf("INSERT INTO authors ( name , country) VALUES ('%s','%s');", author.Name, author.Country)

	result, err := db.Query(sql_query)

	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
}

// Updates an author in database
func db_update_author(id string, author Author) {
	db := db_conn()
	defer db.Close()

	old_author := db_get_author_by_id(id)

	// Keep old author info or update

	if author.Name == "" {
		author.Name = old_author.Name
	} else { // add slashes
		author.Name = strings.ReplaceAll(author.Name, "'", "\\'")
	}
	if author.Country == "" {
		author.Country = old_author.Country
	} else { // add slashes
		author.Name = strings.ReplaceAll(author.Name, "'", "\\'")
	}

	sql_query := fmt.Sprintf("UPDATE authors SET name ='%s',  country ='%s' WHERE id=%s;", author.Name, author.Country, old_author.ID)

	result, err := db.Query(sql_query)

	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
}

// Deletes an author with ID = num
//
// Example: db_delete_author(1) , deletes the author with id = 1
func db_delete_author(num int) {
	db := db_conn()
	defer db.Close()

	sql_query := fmt.Sprintf("DELETE from authors WHERE id=%d;", num)

	result, err := db.Query(sql_query)

	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	// get all books after the deleted book
	sql_query = fmt.Sprintf("SELECT * from authors WHERE id >%d;", num)

	result2, err := db.Query(sql_query)

	if err != nil {
		panic(err.Error())
	}

	defer result2.Close()

	for result2.Next() {
		var author Author
		// reset id for every author after the deleted author
		err = result2.Scan(&author.ID, &author.Name, &author.Country)
		if err != nil {
			panic(err.Error())
		}
		id, err := strconv.Atoi(author.ID)
		if err != nil {
			panic(err.Error())
		}
		var newId int = id - 1
		// set id-1
		sql_query = fmt.Sprintf("UPDATE authors SET id = %d WHERE id=%d;", newId, id)
		result3, err := db.Query(sql_query)
		if err != nil {
			panic(err.Error())
		}
		result3.Close()

	}

	// reset auto_increment
	sql_query = "SELECT COUNT(1) from authors;"
	result4, err := db.Query(sql_query)
	if err != nil {
		panic(err.Error())
	}
	defer result4.Close()
	if result4.Next() {
		var total int
		err = result4.Scan(&total)
		if err != nil {
			panic(err.Error())
		}
		sql_query = fmt.Sprintf("ALTER TABLE authors AUTO_INCREMENT = %d;", total)
		result5, err := db.Query(sql_query)
		if err != nil {
			panic(err.Error())
		}
		defer result5.Close()
	}
}
