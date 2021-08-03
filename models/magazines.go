package models

type Magazine struct {
	title string
	isbn string
	authors []string
	publishedAt string
}

func (magazine Magazine) MapToLiteraryItem() LiteraryItem{
	return LiteraryItem{
		title:   magazine.title,
		isbn:    magazine.isbn,
		authors: magazine.authors,
		other:   magazine.publishedAt,
	}
}