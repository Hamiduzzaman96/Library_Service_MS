package repository

import (
	"context"

	"github.com/Hamiduzzaman96/Library_Service_MS/internal/domain"
)

type LoanRepository interface {
	Create(ctx context.Context, loan *domain.Loan) error
	GetByID(ctx context.Context, id int64) (*domain.Loan, error)
}
