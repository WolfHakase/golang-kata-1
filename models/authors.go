package models

import "fmt"

type Authors []Author

type Author struct {
	email     string
	firstname string
	lastname  string
}

func (author Author) Print() {
	fmt.Printf("%s, %s, %s\n", author.email, author.firstname, author.lastname)
}
