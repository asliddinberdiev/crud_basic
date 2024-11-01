package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/asliddinberdiev/crud_basic/internal/domain"
)

type Books struct {
	db *sql.DB
}

func NewBooks(db *sql.DB) *Books {
	return &Books{db: db}
}

func (r *Books) Create(ctx context.Context, book domain.Book) error {
	_, err := r.db.Exec("INSERT INTO books (title, author, publish_date, rating) values($1, $2, $3, $4)", book.Title, book.Author, book.PublishDate, book.Rating)

	return err
}

func (r *Books) GetByID(ctx context.Context, id int64) (domain.Book, error) {
	var book domain.Book
	err := r.db.QueryRow("SELECT id, title, author, publish_date, rating FROM books WHERE id = $1", id).Scan(&book.ID, &book.Title, &book.Author, &book.PublishDate, &book.Rating)
	if err == sql.ErrNoRows {
		return book, domain.ErrBookNotFound
	}

	return book, err
}

func (r *Books) GetAll(ctx context.Context) ([]domain.Book, error) {
	rows, err := r.db.Query("SELECT id, title, author, publish_date, rating FROM books")
	if err != nil {
		return nil, err
	}

	books := make([]domain.Book, 0)
	for rows.Next() {
		var book domain.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublishDate, &book.Rating); err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, rows.Err()
}

func (r *Books) Delete(ctx context.Context, id int64) error {
	_, err := r.db.Exec("DELETE FROM books WHERE id = $1", id)

	return err
}

func (r *Books) Update(ctx context.Context, id int64, input domain.UpdateBookInput) error {
	setValues := make([]string, 0)
	args := make([]any, 0)
	argID := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title = $%d", argID))
		args = append(args, *input.Title)
		argID++
	}

	if input.Author != nil {
		setValues = append(setValues, fmt.Sprintf("author = $%d", argID))
		args = append(args, *input.Author)
		argID++
	}

	if input.PublishDate != nil {
		setValues = append(setValues, fmt.Sprintf("publish_date = $%d", argID))
		args = append(args, *input.PublishDate)
		argID++
	}

	if input.Rating != nil {
		setValues = append(setValues, fmt.Sprintf("rating = $%d", argID))
		args = append(args, *input.Rating)
		argID++
	}

	_, err := r.db.Exec("UPDATE books SET title = $1, author = $2, publish_date = $3, rating = $4 WHERE id = $5", input.Title, input.Author, input.PublishDate, input.Rating, id)
	if err == sql.ErrNoRows {
		return domain.ErrBookNotFound
	}

	return err
}
