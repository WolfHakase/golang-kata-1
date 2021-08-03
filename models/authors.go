package models

import "fmt"

type Authors []Author

type Author struct {
	email     string
	firstname string
	lastname  string
}

func (author *Author) SetValues(email, firstname, lastname string) {
 	author.email = email
 	author.firstname = firstname
 	author.lastname = lastname
}

func (author Author) Print() {
	fmt.Printf("%s, %s, %s\n", author.email, author.firstname, author.lastname)
}
