package models

import (
	"errors"
	"fmt"
	"github.com/echocat/golang-kata-1/v1/helpers"
	"strings"
)

type Books []Book

func (books Books) FindByISBN(isbn string) (book *Book, err error){
	for _, book := range books {
		if book.isbn == isbn{
			return &book, nil
		}
	}
	return nil, errors.New("book with ISBN not found")
}

func (books Books) FindByAuthorsEmail(authorEmail string) (book *Book, err error) {
	for _, book := range books {
		if helpers.Contains(book.authors, authorEmail) {
			return &book, nil
		}
	}
	return nil, errors.New("book with author not found")
}

type Book struct {
	title string
	isbn string
	authors []string
	description string
}

func (book Book) Print() {
	authorsEmails := strings.Join(book.authors, ",")
	fmt.Printf("%s, %s, %s, %s\n", book.title, book.isbn, authorsEmails, book.description)
}
