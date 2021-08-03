package models

type Magazines []Magazine

func (magazines Magazines) MapToLiteraryItems() LiteraryItems {
	items := LiteraryItems{}
	for _, magazine := range magazines{
		item := magazine.MapToLiteraryItem()
		items = append(items, item)
	}
	return items
}

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