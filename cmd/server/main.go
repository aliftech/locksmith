package main

import (
	"log"
	"net"
	"os"

	// 👈 Import repo
	"github.com/aliftech/locksmith/internal/core/app/repositories/mysql"
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
	// Initialize repository with DB connection
	walletRepo := mysql.NewWalletRepository(config.ConnectDB())

	// Create gRPC server
	lis, err := net.Listen("tcp", os.Getenv("GRPC_TCP"))
	if err != nil {
		log.Fatal("failed to listen:", err)
	}

	s := grpc.NewServer()
	// ✅ Pass repository interface, not *sql.DB
	walletpb.RegisterWalletServiceServer(s, server.NewWalletServer(walletRepo))

	log.Println("gRPC server listening on", os.Getenv("GRPC_TCP"))
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve:", err)
	}
}
