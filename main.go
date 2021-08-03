package main

import (
	"errors"
	"fmt"
	"github.com/echocat/golang-kata-1/v1/files"
	"github.com/echocat/golang-kata-1/v1/models"
)

var literaryItems models.LiteraryItems = nil
var authors models.Authors = nil

func main() {
	fmt.Println(welcomeMessage())

	readAuthors()
	readLiteraryItems()

	fmt.Println("Print items:")
	err := printItems()
	if err != nil {
		fmt.Printf("Could not print items, because: %s\n", err.Error())
	}

	isbnSlice := []string{"4545-8558-3232", "2215-0012-5487", "1313-4545-8875"}
	for _, isbn := range isbnSlice {
		fmt.Printf("Print a book with isbn %s:\n", isbn)
		err = printItemForISBN(isbn)
		if err != nil {
			fmt.Printf("Could not find book with isbn, because: %s\n", err.Error())
		}
	}

	authorsEmails := []string{"null-ferdinand@echocat.org", "null-rabe@echocat.org", "null-walter@echocat.org"}
	for _, email := range authorsEmails {
		fmt.Printf("Print a book with author %s:\n", email)
		err = printItemsForAuthor(email)
		if err != nil {
			fmt.Printf("Could not find book with author, because: %s\n", err.Error())
		}
	}

	fmt.Println("Print all books and magazines sorted by title:")
	err = printItemsSortedByTitle()
	if err != nil {
		fmt.Printf("Could not print items sorted, because: %s\n", err.Error())
	}
}

func readAuthors() {
	csvAuthors := files.CSVFile{Name: "resources/authors.csv"}
	authorsFromFile, err := csvAuthors.GetAuthors()

	if err != nil {
		fmt.Printf("Could not read authors, because: %s\n", err.Error())
		return
	}

	authors = *authorsFromFile
}

func readLiteraryItems() {
	csvBooks := files.CSVFile{Name: "resources/books.csv"}
	booksFromFile, err := csvBooks.GetLiteraryItems()

	if err != nil {
		fmt.Printf("Could not read books, because: %s\n", err.Error())
		return
	}

	csvMagazine := files.CSVFile{Name: "resources/magazines.csv"}
	magazinesFromFile, err := csvMagazine.GetLiteraryItems()

	if err != nil {
		fmt.Printf("Could not read magazines, because: %s\n", err.Error())
		return
	}

	literaryItems = append(*booksFromFile, *magazinesFromFile...)
}

func isInstantiatedLiteraryItems() error {
	if literaryItems == nil {
		return errors.New("no literary items found")
	}
	return nil
}

func printItemForISBN(isbn string) error {
	err := isInstantiatedLiteraryItems()
	if err != nil {
		return err
	}

	item, err := literaryItems.FindByISBN(isbn)
	if err != nil {
		return err
	}

	item.Print()
	return nil
}

func printItemsForAuthor(authorEmail string) error {
	err := isInstantiatedLiteraryItems()
	if err != nil {
		return err
	}

	items, err := literaryItems.FindByAuthorsEmail(authorEmail)
	if err != nil {
		return nil
	}

	items.PrintTableHeader()
	items.Print()
	return nil
}

func printItems() error {
	err := isInstantiatedLiteraryItems()
	if err != nil {
		return err
	}

	literaryItems.Print()
	return nil
}

func printItemsSortedByTitle() error {
	err := isInstantiatedLiteraryItems()
	if err != nil {
		return err
	}

	literaryItems.PrintSortedByTitle()
	return nil
}

func welcomeMessage() string {
	return "Hello world!"
}
