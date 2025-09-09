package server

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"fmt"

	walletpb "github.com/aliftech/locksmith/internal/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WalletServer struct {
	walletpb.UnimplementedWalletServiceServer
	DB *sql.DB
}

func NewWalletServer(db *sql.DB) *WalletServer {
	return &WalletServer{DB: db}
}

func (s *WalletServer) StoreWallet(ctx context.Context, req *walletpb.StoreWalletRequest) (*walletpb.StoreWalletResponse, error) {
	// Hash passphrase for storage (never store plain text)
	passHash := fmt.Sprintf("%x", sha256.Sum256([]byte(req.PassphraseHash)))

	// Insert into MySQL
	query := `
		INSERT INTO wallets (mnemonic, public_key, private_key, address, derivation_index, passphrase_hash, created_by, created_at, updated_at, deleted_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, NOW(), NOW(), NULL)
	`

	result, err := s.DB.ExecContext(ctx, query,
		req.Mnemonic,
		req.PublicKey,
		req.PrivateKey,
		req.Address,
		req.Index,
		passHash,
		"CLI system",
	)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to store wallet: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get wallet ID: %v", err)
	}

	return &walletpb.StoreWalletResponse{
		Success:  true,
		Message:  "Wallet stored successfully",
		WalletId: id,
	}, nil
}
