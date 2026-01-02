package grpct

import (
	"context"

	"github.com/Hamiduzzaman96/Library_Service_MS/internal/domain"
	"github.com/Hamiduzzaman96/Library_Service_MS/internal/usecase"
	"github.com/Hamiduzzaman96/Library_Service_MS/proto/loanpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LoanHandler struct {
	loanpb.UnimplementedLoanServiceServer
	usecase *usecase.LoanUsecase
}

func NewLoanHandler(u *usecase.LoanUsecase) *LoanHandler {
	return &LoanHandler{usecase: u}
}

func (h *LoanHandler) CreateLoan(ctx context.Context, req *loanpb.CreateLoanRequest) (*loanpb.CreateLoanResponse, error) {
	// proto -> domain
	loan := &domain.Loan{
		BookID: req.BookId,
		UserID: req.UserId,
	}
	// call usecase
	err := h.usecase.CreateLoan(ctx, loan)
	if err != nil {
		//business error mapping
		if err == usecase.ErrBookNotFound {
			return nil, status.Errorf(codes.NotFound, "book not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to create loan : %v", err)
	}
	//domain -> proto
	return &loanpb.CreateLoanResponse{
		Id:     loan.ID,
		BookId: loan.BookID,
		UserId: loan.UserID,
		Status: loan.Status,
	}, nil
}

// Get loan

func (h *LoanHandler) GetLoan(ctx context.Context, req *loanpb.GetLoanRequest) (*loanpb.GetLoanResponse, error) {
	loan, err := h.usecase.GetLoan(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get loan: %v", err)
	}
	if loan == nil {
		return nil, status.Errorf(codes.NotFound, "loan not found")
	}
	return &loanpb.GetLoanResponse{
		Id:     loan.ID,
		BookId: loan.BookID,
		UserId: loan.UserID,
		Status: loan.Status,
	}, nil
}
