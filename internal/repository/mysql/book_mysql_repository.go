package mysql

import (
	"context"
	"database/sql"

	"github.com/Hamiduzzaman96/Library_Service_MS/internal/domain"
)

type bookMySQLRepository struct {
	db *sql.DB
}

func NewBookMySQLRepository(db *sql.DB) *bookMySQLRepository {
	return &bookMySQLRepository{db: db}
}

func (r *bookMySQLRepository) Create(ctx context.Context, book *domain.Book) error {
	query := `INSERT INTO books (title, author) VALUES (?,?)`

	result, err := r.db.ExecContext(ctx, query, book.Title, book.Author)

	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	book.ID = id
	return nil
}

func (r *bookMySQLRepository) GetByID(ctx context.Context, id int64) (*domain.Book, error) {
	query := `SELECT id,title, author FROM books WHERE id = ?`
	book := &domain.Book{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&book.ID, &book.Title, &book.Author,
	)
	if err != nil {
		return nil, err
	}
	return book, nil

}
