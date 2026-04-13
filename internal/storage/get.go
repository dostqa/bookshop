package storage

import (
	"bookshop/internal/models"
	"fmt"
)

func (s *Storage) GetAllBooks() ([]models.Book, error) {
	const op = "storage.GetAll"
	var res []models.Book

	rows, err := s.db.Query(`
    SELECT 
        books.id,
        books.title,
        categories.category_name,
        authors.first_name || ' ' || authors.second_name AS author
    FROM books
    LEFT JOIN authors ON books.author = authors.id
    LEFT JOIN categories ON books.category = categories.id;
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var book models.Book

		err = rows.Scan(&book.ID, &book.Title, &book.Category, &book.Author)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		res = append(res, book)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}

	return res, nil
}
