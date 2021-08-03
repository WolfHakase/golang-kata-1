package main

import (
	"errors"
	"fmt"
	"github.com/echocat/golang-kata-1/v1/models"
)

var literaryItems models.LiteraryItems = nil

func main() {
	fmt.Println(welcomeMessage())
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
