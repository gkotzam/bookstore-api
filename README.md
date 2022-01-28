# bookstore-api



# Usage

## Get all Books
```
[GET] /api/books
```

## Get a Book by {id}
```
[GET] /api/books/{id}
```

## Create a Book
```
[POST] /api/books

body:
{
    "title": "title",
    "isbn10": "isbn_num",
    "isbn13": "isbn13_num",
    "authors": [{"name": "Author1"},{"name": "Autho2"},{"name": "Autho3"}],
    "publisher": "publisher",
    "pub_year": "0000",
    "language": "English",
    "pages": "000",
    "description": "description...",
    "price": "00.00"
    }

```








