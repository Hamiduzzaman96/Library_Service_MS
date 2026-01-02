package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/Hamiduzzaman96/Library_Service_MS/internal/handler/grpct"
	"github.com/Hamiduzzaman96/Library_Service_MS/internal/repository/mysql"
	"github.com/Hamiduzzaman96/Library_Service_MS/internal/usecase"
	"github.com/Hamiduzzaman96/Library_Service_MS/pkg/config"
	"github.com/Hamiduzzaman96/Library_Service_MS/proto/userpb"
	"google.golang.org/grpc"
)

func main() {
	//connect to MySQL
	// db, err := sql.Open("mysql", "root:pass @tcp(localhost:3306)/library")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if err := db.Ping(); err != nil {
	// 	log.Fatalf("failed to connect databse: %v", err)
	// }
	// defer db.Close()

	cfg := config.LoadUserConfig()

	db, err := sql.Open("mysql", cfg.DBDSN)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to connect databse: %v", err)
	}
	defer db.Close()

	repo := mysql.NewUserMySQLRepository(db)
	u := usecase.NewUserUsecase(repo)

	// grpc server
	handler := grpct.NewUserHandler(u)
	lis, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		log.Fatalf("failed to listen gRPC : %v", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, handler)

	go func() {
		log.Println("gRPC user service running on: %s", cfg.GRPCPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("gRPC server error :%v", err)
		}
	}()
}
