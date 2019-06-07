package services

import (
	"api/models"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

type IBookService interface {
	CreateBook(book models.Book) (int, error)
	RetrieveBookById(id int) (models.Book, error)
	RetrieveAllBooks() ([]models.Book, error)
	UpdateBookById(book models.Book, id int) error
	DeleteBookById(id int) error
}

func CreateBook(book models.Book) (int, error) {
	db, err := GetDatabase()
	if err != nil {
		return 42, err
	}
	res, err := db.Exec(
		`INSERT INTO library (title, author, publisher, publish_date, rating, status) VALUES (?, ?, ?, ?, ?, ?)`,
		book.Title,
		book.Author,
		book.Publisher,
		book.PublishDate,
		book.Rating,
		book.Status,
	)
	if err != nil {
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
		`SELECT title, author, publisher, publish_date, rating, status FROM library WHERE id = ?`,
		id,
	)
	if err = row.Scan(
		&book.Title,
		&book.Author,
		&book.Publisher,
		&book.PublishDate,
		&book.Rating,
		&book.Status,
	); err != nil {
		return models.Book{}, err
	}

	db.Close()
	
	return book, nil
}

func RetrieveAllBooks() ([]models.Book, error) {
	db, err := GetDatabase()
	if err != nil {
		return nil, err
	}

	var books []models.Book

	rows, err := db.Query(
		`SELECT title, author, publisher, publish_date, rating, status FROM library`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var book models.Book
		if err = rows.Scan(
			&book.Title,
			&book.Author,
			&book.Publisher,
			&book.PublishDate,
			&book.Rating,
			&book.Status,
		); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.New("Row iteration error")
	}

	db.Close()

	return books, nil
}

func UpdateBookById(book models.Book, id int) error {
	db, err := GetDatabase()
	if err != nil {
		return err
	}

	res, err := db.Exec(
		`UPDATE library SET title = ?, author = ?, publisher = ?, publish_date = ?, rating = ?, status = ? WHERE id = ?`,
		book.Title,
		book.Author,
		book.Publisher,
		book.PublishDate,
		book.Rating,
		book.Status,
		id,
	)
	if err != nil {
		return err
	}

	db.Close()

	rowsAffected, err := res.RowsAffected()
	if rowsAffected < 1 || err != nil {
		return errors.New("Book not found")
	}

	return nil
}

func DeleteBookById(id int) error {
	db, err := GetDatabase()
	if err != nil {
		return err
	}

	res, err := db.Exec(
		`DELETE FROM library WHERE id = ?`,
		id,
	)
	if err != nil {
		return err
	}

	db.Close()
	
	rowsAffected, err := res.RowsAffected()
	if rowsAffected < 1 || err != nil {
		return errors.New("Book not found")
	}

	return nil
}