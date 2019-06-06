package services

import (
	"api/models"
	"errors"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type IBookService interface {
	CreateBook(book models.Book) (int, error)
	RetrieveBookById(id int) (models.Book, error)
	RetrieveAllBooks() ([]models.Book, error)
	UpdateBookById(book models.Book, id int) error
	DeleteBookById(id int) error
}

func CreateBook(book models.Book) (int, error) {
	tx, err := StartTransaction()
	if err != nil {
		return 42, err
	}
	res, err := tx.Exec(
		`INSERT INTO library VALUES ($1, $2, $3, $4, $5, $6)`,
		book.Title,
		book.Author,
		book.Publisher,
		book.PublishDate,
		book.Rating,
		book.Status,
	)
	if err = CompleteTransaction(tx, err); err != nil {
		return 42, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 42, errors.New("Error extracting book id")
	}
	return int(id), nil
}

func RetrieveBookById(id int) (models.Book, error) {
	db, err := GetDatabase()
	if err != nil {
		return models.Book{}, err
	}

	var book models.Book

	row := db.QueryRow(
		`SELECT title, author, publisher, publish_date, rating, status FROM library WHERE id = $1`,
		id,
	)
	if err = row.Scan(
		book.Title,
		book.Author,
		book.Publisher,
		book.PublishDate,
		book.Rating,
		book.Status,
	); err != nil {
		return models.Book{}, err
	}
	
	return book, nil
}

func RetrieveAllBooks() ([]models.Book, error) {
	db, err := GetDatabase()
	if err != nil {
		return nil, err
	}

	var books []models.Book

	rows, err := db.Query(
		`SELECT title, author, publisher, publish_date, rating, status FROM library`
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var book models.Book
		if err = rows.Scan(
			book.Title,
			book.Author,
			book.Publisher,
			book.PublishDate,
			book.Rating,
			book.Status,
		); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.New("Row iteration error")
	}

	return books, nil
}

func UpdateBookById(book models.Book, id int) error {
	tx, err := StartTransaction()
	if err != nil {
		return err
	}

	res, err := tx.Exec(
		`UPDATE library SET title = $2, author = $3, publisher = $4, publish_date = $5, rating = $6, status = $7 WHERE id = $1`,
		id,
		book.Title,
		book.Author,
		book.Publisher,
		book.PublishDate,
		book.Rating,
		book.Status,
	)
	if err = CompleteTransaction(tx, err); err != nil {
		return err
	}

	return nil
}

func DeleteBookById(id int) error {
	tx, err := StartTransaction()
	if err != nil {
		return err
	}

	res, err := tx.Exec(
		`DELETE FROM library WHERE id = $1`,
		id,
	)
	if err = CompleteTransaction(tx, err); err != nil {
		return err
	}

	return nil
}