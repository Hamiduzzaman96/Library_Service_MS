package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/Hamiduzzaman96/Library_Service_MS/internal/domain"
	"github.com/Hamiduzzaman96/Library_Service_MS/internal/handler/grpct"
	"github.com/Hamiduzzaman96/Library_Service_MS/internal/repository/mysql"
	"github.com/Hamiduzzaman96/Library_Service_MS/internal/usecase"
	"github.com/Hamiduzzaman96/Library_Service_MS/proto/bookpb"
	"github.com/Hamiduzzaman96/Library_Service_MS/proto/loanpb"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/user_service")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to connect databse: %v", err)
	}
	defer db.Close()

	loanRepo := mysql.NewLoanMySQLRepository(db)
	bookConn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect book service: %v", err)
	}
	defer bookConn.Close()
	bookGRPC := bookpb.NewBookServiceClient(bookConn)
	bookAdapter := domain.NewBookGRPCClient(bookGRPC)

	loanUsecase := usecase.NewLoanUsecase(loanRepo, bookAdapter)

	loanHandler := grpct.NewLoanHandler(loanUsecase)

	grpcServer := grpc.NewServer()
	loanpb.RegisterLoanServiceServer(grpcServer, loanHandler)

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Loan Service running on :50053")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
