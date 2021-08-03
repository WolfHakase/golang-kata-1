package models

import (
	"errors"
	"fmt"
	"github.com/echocat/golang-kata-1/v1/helpers"
	"sort"
	"strings"
	"unicode/utf8"
)

type LiteraryItems []LiteraryItem

func (items LiteraryItems) Len() int {
	return len(items)
}

func (items LiteraryItems) Less(i, j int) bool {
	iRune, _ := utf8.DecodeRuneInString(items[i].title)
	jRune, _ := utf8.DecodeRuneInString(items[j].title)
	return iRune < jRune
}

func (items LiteraryItems) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

func (items LiteraryItems) PrintSortedByTitle() {
	sort.Sort(items)
	fmt.Print("Title, ISBN, Authors, Other info")
	for _, item := range items {
		item.Print()
	}
}

func (items LiteraryItems) FindByISBN(isbn string) (item *LiteraryItem, err error) {
	for _, item := range items {
		if item.isbn == isbn {
			return &item, nil
		}
	}
	return nil, errors.New("literary item with ISBN not found")
}

func (items LiteraryItems) FindByAuthorsEmail(authorEmail string) (item *LiteraryItem, err error) {
	for _, item := range items {
		if helpers.Contains(item.authors, authorEmail) {
			return &item, nil
		}
	}
	return nil, errors.New("literary item with author not found")
}

type LiteraryItem struct {
	title   string
	isbn    string
	authors []string
	other   string
}

func (item LiteraryItem) Print() {
	authorsEmails := strings.Join(item.authors, ",")
	fmt.Printf("%s, %s, [%s], %s\n", item.title, item.isbn, authorsEmails, item.other)
}
