package models

type Books []Book

func (books Books) MapToLiteraryItems() LiteraryItems {
	items := LiteraryItems{}
	for _, book := range books{
		item := book.MapToLiteraryItem()
		items = append(items, item)
	}
	return items
}

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

