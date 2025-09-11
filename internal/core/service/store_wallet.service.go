package service

import (
	"context"
	"fmt"
	"os"
	"time"

	walletpb "github.com/aliftech/locksmith/internal/grpc"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/aliftech/locksmith/internal/core/lib"
	"github.com/aliftech/locksmith/internal/core/util"
)

func StoreWalletViaGRPC(mnemonic string, ticker string, cryptoWallet *util.CryptoAddress, index uint32, passphrase string) error {
	// Connect to gRPC server
	conn, err := grpc.NewClient(os.Getenv("GRPC_SERVER"), grpc.WithTransportCredentials(insecure.NewCredentials()))
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

	encryptedPrivateKey, enPrivKeyErr := bcrypt.GenerateFromPassword([]byte(cryptoWallet.PrivateKeyHex), 14)
	if enPrivKeyErr != nil {
		fmt.Println(lib.Red("Encryption ERROR: ", enPrivKeyErr))
		return fmt.Errorf("encryption error: %s", enPrivKeyErr)
	}

	// Send request
	req := &walletpb.StoreWalletRequest{
		Mnemonic:       mnemonic,
		Ticker:         ticker,
		PublicKey:      cryptoWallet.PublicKeyHex,
		PrivateKey:     string(encryptedPrivateKey),
		Address:        cryptoWallet.Address,
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

func StoreBTCWalletViaGRPC(mnemonic string, ticker string, btcWallet *util.BitcoinAddress, index uint32, passphrase string, purpose uint32) error {
	// connect to gRPC server
	conn, err := grpc.NewClient(os.Getenv("GRPC_SERVER"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}

	defer conn.Close()

	client := walletpb.NewWalletServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// encrypt private key and passphrase
	encryptedPassphrase, passphraseErr := bcrypt.GenerateFromPassword([]byte(passphrase), 14)
	if passphraseErr != nil {
		fmt.Println(lib.Red("Encryption ERROR: ", passphraseErr))
		return fmt.Errorf("encryption error: %s", passphraseErr)
	}

	encryptedPrivateKey, enPrivKeyErr := bcrypt.GenerateFromPassword([]byte(btcWallet.PrivateKeyHex), 14)
	if enPrivKeyErr != nil {
		fmt.Println(lib.Red("Encryption ERROR: ", enPrivKeyErr))
		return fmt.Errorf("encryption error: %s", enPrivKeyErr)
	}

	address := ""

	switch {
	case purpose == 0x8000002C:
		address = btcWallet.P2PKHAddress
	case purpose == 0x80000054:
		address = btcWallet.P2WPKHAddress
	case purpose == 0x80000056:
		address = btcWallet.P2TRAddress
	default:
		address = btcWallet.WIF
	}

	req := &walletpb.StoreWalletRequest{
		Mnemonic:       mnemonic,
		Ticker:         ticker,
		PublicKey:      btcWallet.PublicKeyHex,
		PrivateKey:     string(encryptedPrivateKey),
		Address:        address,
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
