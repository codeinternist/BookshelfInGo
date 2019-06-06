package services

import (
	"errors"
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", "reader:secret@localhost:3306/library")
	if err != nil {
		return nil, errors.New("Database failed to open")
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, errors.New("Database connection failure")
	}

	return db, nil
}

func StartTransaction() (*sql.Tx, error) {
	db, err := GetDatabase()
	if err != nil {
		return nil, err
	}

	var ctx context.Context

	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, errors.New("Transaction initiation failure")
	}

	return tx, nil
}

func CompleteTransaction(tx *sql.Tx, err error) error {
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}