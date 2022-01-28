# bookstore-api



# Usage
## BOOKS
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

## Update a Book{id}
```
[PATCH] /api/books/{id}

example: 
body:
{
    "title": "updated title",
    "isbn10": "updated isbn",
    "isbn13": "updated inbn13",
    "authors": [{"name": "updated author1"},{"name": "updated author2"}],
    "publisher": "updated publisher",
    "price": "00.00"
}

```

## Delete a Book{id}
```
[DELETE] /api/books/{id}
```

## AUTHORS
## Get all Authors
```
[GET] /api/authors
```

## Get an Author by {id}
```
[GET] /api/authors/{id}
```

## Create an Author
```
[POST] /api/authors

body:
{
    "name": "Author",
    "country": "Country"
}
```

## Update an Author{id}
```
[PATCH] /api/authors/{id}

example: 
body:
{
    "country": "new Country"
}

```

## Delete an Author{id}
```
[DELETE] /api/authors/{id}
```
## SEARCH
## Search a Book by title
```
[GET] /api/search/books-by-title/{title}
```
## Search a Book by publisher
```
[GET] /api/search/books-by-publisher/{publisher}
```
## Search a Book by Author
```
[GET] /api/search/books-by-author/{author}
```





