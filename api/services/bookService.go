package services

import (
	"api/models"
	// database service
)

type IBookService interface {
	CreateBook(book Book) (int, error)
	RetrieveBookById(id int) (Book, error)
	RetrieveAllBooks() ([]Book, error)
	UpdateBookById(book Book, id int) (error)
	DeleteBookById(id int) (error)
}