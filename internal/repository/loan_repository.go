package repository

import "github.com/Hamiduzzaman96/Library_Service_MS/internal/domain"

type LoanRepository interface {
	Create(loan *domain.Loan) error
	GetByID(id int64) (*domain.Loan, error)
}
