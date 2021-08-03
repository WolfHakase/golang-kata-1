package models

type Books []Book

type Book struct {
	title string
	isbn string
	authors []string
	description string
}
