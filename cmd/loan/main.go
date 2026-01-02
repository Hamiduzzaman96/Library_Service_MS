package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/Hamiduzzaman96/Library_Service_MS/internal/handler/grpct"
	"github.com/Hamiduzzaman96/Library_Service_MS/internal/repository/mysql"
	"github.com/Hamiduzzaman96/Library_Service_MS/internal/usecase"
	"github.com/Hamiduzzaman96/Library_Service_MS/pkg/config"
	"github.com/Hamiduzzaman96/Library_Service_MS/proto/bookpb"
	"github.com/Hamiduzzaman96/Library_Service_MS/proto/loanpb"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.LoadLoanConfig()
	bookConn, err := grpc.Dial(cfg.BookServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect Book service: %v", err)
	}
	defer bookConn.Close()

	db, err := sql.Open("mysql", cfg.DBDSN)

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to connect databse: %v", err)
	}
	defer db.Close()

	loanRepo := mysql.NewLoanMySQLRepository(db)
	bookClient := bookpb.NewBookServiceClient(bookConn)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	loanUsecase := usecase.NewLoanUsecase(loanRepo, bookClient)

	loanHandler := grpct.NewLoanHandler(loanUsecase)

	grpcServer := grpc.NewServer()
	loanpb.RegisterLoanServiceServer(grpcServer, loanHandler)

	lis, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		log.Fatalf("failed to listen gRPC: %v", err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
