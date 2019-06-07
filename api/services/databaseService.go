package services

import (
	"errors"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", "reader:secret@tcp(localhost:3306)/library?parseTime=true")
	if err != nil {
		return nil, errors.New("Database failed to open")
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.New("Database connection failure")
	}

	return db, nil
}