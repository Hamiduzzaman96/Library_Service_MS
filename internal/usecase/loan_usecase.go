package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/Hamiduzzaman96/Library_Service_MS/internal/domain"
	"github.com/Hamiduzzaman96/Library_Service_MS/internal/repository"
	"github.com/Hamiduzzaman96/Library_Service_MS/pkg/resilience"
	"github.com/sony/gobreaker"
)

var ErrBookNotFound = errors.New("Book Not Found")

type LoanUsecase struct {
	loanRepo repository.LoanRepository
	bookSvc  domain.BookServiceClient
	cb       *gobreaker.CircuitBreaker
	timeout  time.Duration
	retries  int
	delay    time.Duration
}

func NewLoanUsecase(loanRepo repository.LoanRepository, bookSvc domain.BookServiceClient) *LoanUsecase {
	return &LoanUsecase{
		loanRepo: loanRepo,
		bookSvc:  bookSvc,
		cb:       resilience.NewBreaker("loan-create"),
		timeout:  3 * time.Second,
		retries:  3,
		delay:    500 * time.Millisecond,
	}

}

func (u *LoanUsecase) CreateLoan(ctx context.Context, loan *domain.Loan) error {
	//timeout apply

	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	//Retry Wrapper

	return resilience.Retry(u.retries, u.delay, func() error {
		//circuit breaker wrapper
		_, err := u.cb.Execute(func() (interface{}, error) {

			//service - to - service call, book exists?
			exists, err := u.bookSvc.CheckBookExists(ctx, loan.BookID)
			if err != nil {
				return nil, err
			}
			if !exists {
				return nil, ErrBookNotFound
			}

			//set loan initial status
			loan.Status = "CREATED"

			//persist loan
			return nil, u.loanRepo.Create(ctx, loan)
		})
		return err
	})
}

func (u *LoanUsecase) GetLoan(ctx context.Context, id int64) (*domain.Loan, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	return u.loanRepo.GetByID(ctx, id)
}
