package storage

import (
	"database/sql"
	"fmt"

	// подключение драйвера базы данных sqlite
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(storagePath string) (*Storage, error) {
	const op = "storage.NewStorage"

	// _foreign_keys = ON; в SQLite активирует проверку ограничений внешних ключей до соединения
	db, err := sql.Open("sqlite3", storagePath+"?_foreign_keys=on")
	if err != nil {
		return nil, fmt.Errorf("%s:\n %w", op, err)
	}

	// создание таблицы authors
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS authors (
    	id INTEGER PRIMARY KEY,
    	first_name TEXT NOT NULL,
    	second_name TEXT NOT NULL);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// создание таблицы categories
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS categories (
		id INTEGER PRIMARY KEY,
		category_name TEXT NOT NULL UNIQUE);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// создание таблицы books
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS books (
    	id INTEGER PRIMARY KEY,
    	title TEXT NOT NULL,
    	category INTEGER,
    	author INTEGER,
    FOREIGN KEY (author) REFERENCES authors(id),
	FOREIGN KEY (category) REFERENCES categories(id));
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}
