package server

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/aliftech/locksmith/internal/core/app/models"
	"github.com/aliftech/locksmith/internal/core/app/repositories/interfaces"
	walletpb "github.com/aliftech/locksmith/internal/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WalletServer struct {
	walletpb.UnimplementedWalletServiceServer
	WalletRepo interfaces.WalletInterface
}

func NewWalletServer(walletRepo interfaces.WalletInterface) *WalletServer {
	return &WalletServer{
		WalletRepo: walletRepo,
	}
}

func (s *WalletServer) StoreWallet(ctx context.Context, req *walletpb.StoreWalletRequest) (*walletpb.StoreWalletResponse, error) {
	// Hash passphrase for storage (never store plain text)
	passHash := fmt.Sprintf("%x", sha256.Sum256([]byte(req.PassphraseHash)))

	// Insert into MySQL
	wallet := models.Wallet{
		Mnemonic:        req.Mnemonic,
		Ticker:          req.Ticker,
		PublicKey:       req.PublicKey,
		PrivateKey:      req.PrivateKey,
		Address:         req.Address,
		DerivationIndex: uint(req.Index),
		PassphraseHash:  passHash,
		CreatedBy:       "CLI system",
	}

	// Persist user
	if err := s.WalletRepo.StoreWallet(&wallet); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to store wallet: %v", err)
	}

	return &walletpb.StoreWalletResponse{
		Success:  true,
		Message:  "Wallet stored successfully",
		WalletId: int64(wallet.ID),
	}, nil
}
