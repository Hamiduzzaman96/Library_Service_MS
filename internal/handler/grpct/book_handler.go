package grpct

import (
	"context"

	"github.com/Hamiduzzaman96/Library_Service_MS/internal/domain"
	"github.com/Hamiduzzaman96/Library_Service_MS/internal/usecase"
	pb "github.com/Hamiduzzaman96/Library_Service_MS/proto/bookpb"
)

type BookHandler struct {
	pb.UnimplementedBookServiceServer
	usecase *usecase.BookUsecase
}

func NewBookHandler(u *usecase.BookUsecase) *BookHandler {
	return &BookHandler{usecase: u}
}

func (h *BookHandler) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.Empty, error) {
	book := &domain.Book{
		Title:  req.Title,
		Author: req.Author,
	}

	err := h.usecase.CreateBook(ctx, book)

	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (h *BookHandler) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.GetBookResponse, error) {
	book, err := h.usecase.GetBook(ctx, req.Id)

	if err != nil {
		return nil, err
	}
	return &pb.GetBookResponse{
		Id:     book.ID,
		Author: book.Author,
	}, nil
}
