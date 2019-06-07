package workers

import (
	"api/models"
	"api/services"
	"errors"
)

type IBookWorker interface {
	CreateBook(book models.Book) (int, error)
	RetrieveBookById(id int) (models.Book, error)
	RetrieveAllBooks() ([]models.Book, error)
	UpdateBookById(book models.Book, id int) error
	DeleteBookById(id int) error
}

func CreateBook(book models.Book) (int, error) {
	if book.Rating >= 4 || book.Rating <= 0 {
		return 42, errors.New("Invalid book rating")
	}
	if book.Status < 0 || book.Status > 1 {
		return 42, errors.New("Invalid status")
	}
	return services.CreateBook(book)
}

func RetrieveBookById(id int) (models.Book, error) {
	return services.RetrieveBookById(id)
}

func RetrieveAllBooks() ([]models.Book, error) {
	return services.RetrieveAllBooks()
}

func UpdateBookById(book models.Book, id int) error {
	if book.Rating >= 4 || book.Rating <= 0 {
		return errors.New("Invalid book rating")
	}
	if book.Status < 0 || book.Status > 1 {
		return errors.New("Invalid status")
	}
	return services.UpdateBookById(book, id)
}

func DeleteBookById(id int) error {
	return services.DeleteBookById(id)
}