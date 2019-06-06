package workers

import (
	"api/models"
	"fmt"
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
	fmt.Println(book.Title)
	fmt.Println(book.Status)
	return 42, nil
}

func RetrieveBookById(id int) (models.Book, error) {
	fmt.Println(id)
	return models.Book{Title:"title"}, nil
}

func RetrieveAllBooks() ([]models.Book, error) {
	return nil, errors.New("Not implemented")
}

func UpdateBookById(book models.Book, id int) error {
	fmt.Println(book.Title)
	fmt.Println(book.Status)
	fmt.Println(id)
	return nil
}

func DeleteBookById(id int) error {
	fmt.Println(id)
	return nil
}