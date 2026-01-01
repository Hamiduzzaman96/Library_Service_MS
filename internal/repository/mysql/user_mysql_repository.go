package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/Hamiduzzaman96/Library_Service_MS/internal/domain"
)

type UserMySQLRepository struct {
	db *sql.DB
}

func NewUserMySQLRepository(db *sql.DB) *UserMySQLRepository {
	return &UserMySQLRepository{db: db}
}

func (r *UserMySQLRepository) Create(ctx context.Context, user *domain.User) error {
	query := `
		INSERT INTO users (name,email,created_at) VALUES (?,?,?)
	`
	_, err := r.db.ExecContext(ctx, query, user.Name, user.Email, time.Now())

	return err
}

func (r *UserMySQLRepository) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	query := `
		SELECT id, name, email, created_at FROM users WHERE id = ?
	`
	row := r.db.QueryRowContext(ctx, query, id)
	user := &domain.User{}

	err := row.Scan(&user.ID, &user.Email, &user.CreateAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}
