package main

import (
	"database/sql"
	"log"
	"net"
	"os"

	"github.com/aliftech/locksmith/internal/core/config"
	walletpb "github.com/aliftech/locksmith/internal/grpc"
	"github.com/aliftech/locksmith/internal/server"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

func init() {
	config.EnvSetup()
}

func main() {
	// Connect to MySQL
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/locksmith")
	if err != nil {
		log.Fatal("failed to connect to DB: ", err)
	}
	defer db.Close()

	// Test connection
	if err := db.Ping(); err != nil {
		log.Fatal("failed to ping DB: ", err)
	}

	// Create gRPC server
	lis, err := net.Listen("tcp", os.Getenv("GRPC_TCP"))
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}

	s := grpc.NewServer()
	walletpb.RegisterWalletServiceServer(s, server.NewWalletServer(db))

	log.Println("gRPC server listening on ", os.Getenv("GRPC_TCP"))
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve: ", err)
	}
}
