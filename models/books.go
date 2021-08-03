package models

type Books []Book

type Book struct {
	title string
	isbn string
	authors []string
	description string
}

func (book Book) MapToLiteraryItem() LiteraryItem{
	return LiteraryItem{
		title:   book.title,
		isbn:    book.isbn,
		authors: book.authors,
		other:   book.description,
	}
}

