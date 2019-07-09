package bookRepository

import (
	"github.com/jmoiron/sqlx"
	"sinistra/books-api/models"
)

type BookRepository struct{}

func (b BookRepository) GetBooks(db *sqlx.DB, books []models.Book) ([]models.Book, error) {
	err := db.Select(&books, "SELECT * FROM books ORDER BY id ASC")

	if err != nil {
		return []models.Book{}, err
	}

	return books, nil
}

func (b BookRepository) GetBook(db *sqlx.DB, book models.Book, id int) (models.Book, error) {
	err := db.Get(&book, "SELECT * FROM books WHERE id=$1", id)

	return book, err
}

func (b BookRepository) AddBook(db *sqlx.DB, book models.Book) (int, error) {
	err := db.QueryRow("insert into books (title, author, year) values($1, $2, $3) RETURNING id;",
		book.Title, book.Author, book.Year).Scan(&book.ID)

	if err != nil {
		return 0, err
	}

	return book.ID, nil
}

func (b BookRepository) UpdateBook(db *sqlx.DB, book models.Book) (int64, error) {
	sql := "update books set title=$1, author=$2, year=$3 where id=$4 RETURNING id"
	result, err := db.Exec(sql, &book.Title, &book.Author, &book.Year, &book.ID)

	if err != nil {
		return 0, err
	}

	rowsUpdated, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsUpdated, nil
}

func (b BookRepository) RemoveBook(db *sqlx.DB, id int) (int64, error) {
	result, err := db.Exec("delete from books where id = $1", id)

	if err != nil {
		return 0, err
	}

	rowsDeleted, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsDeleted, nil
}
