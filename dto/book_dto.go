package dto

import (
	"jwtdemo/models"
)

type BookDto struct {
	BookName string `json:"book_name"`
}

func ToBookDto(book models.Book) BookDto {
	return BookDto{
		BookName: book.BookName,
	}
}
