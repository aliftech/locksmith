package main

import (
	"fmt"
	"log"

	"github.com/aliftech/locksmith/internal/core/util"
)

func main() {
	// Create a new wallet with a passphrase
	wallet, err := util.NewCryptoWallet("my-secure-passphrase")
	if err != nil {
		log.Fatalf("Failed to create wallet: %v", err)
	}

	// Print the mnemonic phrase
	fmt.Printf("Mnemonic: %s\n\n", wallet.Mnemonic)

	// Generate Ethereum address
	ethAddress, err := wallet.GenerateEthereumAddress(0)
	if err != nil {
		log.Fatalf("Failed to generate Ethereum address: %v", err)
	}
	fmt.Println("Ethereum Address:")
	fmt.Printf("Private Key: %s\n", ethAddress.PrivateKeyHex)
	fmt.Printf("Public Key: %s\n", ethAddress.PublicKeyHex)
	fmt.Printf("Address: %s\n\n", ethAddress.Address)

	// Generate Cardano address
	adaAddress, err := wallet.GenerateCardanoAddress(0)
	if err != nil {
		log.Fatalf("Failed to generate Cardano address: %v", err)
	}
	fmt.Println("Cardano Address:")
	fmt.Printf("Private Key: %s\n", adaAddress.PrivateKeyHex)
	fmt.Printf("Public Key: %s\n", adaAddress.PublicKeyHex)
	fmt.Printf("Address: %s\n\n", adaAddress.Address)

	// Generate random Ethereum address (not derived from wallet)
	randomEthAddress, err := util.GenerateRandomCryptoAddress("ethereum")
	if err != nil {
		log.Fatalf("Failed to generate random Ethereum address: %v", err)
	}
	fmt.Println("Random Ethereum Address:")
	fmt.Printf("Private Key: %s\n", randomEthAddress.PrivateKeyHex)
	fmt.Printf("Public Key: %s\n", randomEthAddress.PublicKeyHex)
	fmt.Printf("Address: %s\n", randomEthAddress.Address)
}
