package service

import (
	"context"
	"fmt"
	"time"

	walletpb "github.com/aliftech/locksmith/internal/grpc"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/aliftech/locksmith/internal/core/lib"
	"github.com/aliftech/locksmith/internal/core/util"
)

func StoreWalletViaGRPC(mnemonic string, ticker string, ethWallet *util.CryptoAddress, index uint32, passphrase string) error {
	// Connect to gRPC server
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}
	defer conn.Close()

	client := walletpb.NewWalletServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Encrypt private key and passphrase
	encryptedPassphrase, passphraseErr := bcrypt.GenerateFromPassword([]byte(passphrase), 14)
	if passphraseErr != nil {
		fmt.Println(lib.Red("Encryption ERROR: ", passphraseErr))
		return fmt.Errorf("encryption error: %s", passphraseErr)
	}

	encryptedPrivateKey, enPrivKeyErr := bcrypt.GenerateFromPassword([]byte(ethWallet.PrivateKeyHex), 14)
	if enPrivKeyErr != nil {
		fmt.Println(lib.Red("Encryption ERROR: ", enPrivKeyErr))
		return fmt.Errorf("encryption error: %s", enPrivKeyErr)
	}

	// Send request
	req := &walletpb.StoreWalletRequest{
		Mnemonic:       mnemonic,
		Ticker:         ticker,
		PublicKey:      ethWallet.PublicKeyHex,
		PrivateKey:     string(encryptedPrivateKey),
		Address:        ethWallet.Address,
		Index:          index,
		PassphraseHash: string(encryptedPassphrase), // Server will hash it
	}

	res, err := client.StoreWallet(ctx, req)
	if err != nil {
		fmt.Println(lib.Red("RPC failed: %w", err))
		return fmt.Errorf("RPC failed: %w", err)
	}

	if !res.Success {
		fmt.Println(lib.Red("server error: %s", res.Message))
		return fmt.Errorf("server error: %s", res.Message)
	}

	return nil
}
