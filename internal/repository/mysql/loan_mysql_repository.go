package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/Hamiduzzaman96/Library_Service_MS/internal/domain"
)

type LoanMySQLRepository struct {
	db *sql.DB
}

func NewLoanMySQLRepository(db *sql.DB) *LoanMySQLRepository {
	return &LoanMySQLRepository{db: db}
}

func (r *LoanMySQLRepository) Create(ctx context.Context, loan *domain.Loan) error {
	query := `
		INSERT INTO loans (book_id,user_id,status,created_at) VALUES (?,?,?,?)
	`
	_, err := r.db.ExecContext(ctx, query, loan.BookID, loan.UserID, loan.Status, time.Now())

	if err != nil {
		return err
	}

	return nil
}

func (r *LoanMySQLRepository) GetByID(ctx context.Context, id int64) (*domain.Loan, error) {
	query := `
		SELECT id,book_id, user_id, status,created_at FROM loans WHERE id =?
	`
	row := r.db.QueryRowContext(ctx, query, id)

	loan := &domain.Loan{}

	err := row.Scan(
		&loan.ID,
		&loan.BookID,
		&loan.UserID,
		&loan.Status,
		&loan.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return loan, err
}
