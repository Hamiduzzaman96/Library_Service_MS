package config

import "os"

//Book Service Config
type BookConfig struct {
	DBDSN    string
	GRPCPort string
	HTTPPort string
}

//Loan Service Config
type LoanConfig struct {
	DBDSN           string
	GRPCPort        string
	BookServiceAddr string
	UserServiceAdrr string
}

//User Service Config
type UserConfig struct {
	DBDSN    string
	GRPCPort string
}

func LoadBookConfig() BookConfig { //constructor
	return BookConfig{
		DBDSN:    os.Getenv("BOOK_DB_DSN"),
		GRPCPort: os.Getenv("BOOK_GRPC_PORT"),
		HTTPPort: os.Getenv("BOOK_HTTP_PORT"),
	}
}

func LoadLoanConfig() LoanConfig { //constructor
	return LoanConfig{
		DBDSN:           os.Getenv("LOAN_DB_DSN"),
		GRPCPort:        os.Getenv("LOAN_GRPC_PORT"),
		BookServiceAddr: os.Getenv("BOOK_SERVICE_GRPC_ADDR"),
		UserServiceAdrr: os.Getenv("USER_SERVICE_GRPC_ADDR"),
	}
}

func LoadUserConfig() UserConfig { //constructor
	return UserConfig{
		DBDSN:    os.Getenv("USER_DB_DSN"),
		GRPCPort: os.Getenv("USER_GRPC_PORT"),
	}
}
