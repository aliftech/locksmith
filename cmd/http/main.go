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
	fmt.Printf("🔐 Mnemonic Seed Phrase:\n%s\n\n", wallet.Mnemonic)

	// ========================
	// 🟢 HD DERIVED ADDRESSES (from mnemonic)
	// ========================

	fmt.Println("===========================================")
	fmt.Println("🟢 HD DERIVED ADDRESSES (from mnemonic)")
	fmt.Println("===========================================")

	generateAndPrintHDAddress(wallet, "Ethereum", wallet.GenerateEthereumAddress, 0)
	generateAndPrintHDAddress(wallet, "Cardano", wallet.GenerateCardanoAddress, 0)
	generateAndPrintHDAddress(wallet, "Litecoin", wallet.GenerateLitecoinAddress, 0)
	generateAndPrintHDAddress(wallet, "Dogecoin", wallet.GenerateDogecoinAddress, 0)
	generateAndPrintHDAddress(wallet, "Bitcoin Cash", wallet.GenerateBitcoinCashAddress, 0)
	generateAndPrintHDAddress(wallet, "Tron", wallet.GenerateTronAddress, 0)
	generateAndPrintHDAddress(wallet, "Polkadot", wallet.GeneratePolkadotAddress, 0)
	generateAndPrintHDAddress(wallet, "Ripple", wallet.GenerateRippleAddress, 0)

	// ========================
	// 🎲 RANDOMLY GENERATED ADDRESSES (not from wallet)
	// ========================

	fmt.Println("\n===========================================")
	fmt.Println("🎲 RANDOMLY GENERATED ADDRESSES")
	fmt.Println("===========================================")

	generateAndPrintRandomAddress("Bitcoin", "bitcoin")
	generateAndPrintRandomAddress("Ethereum", "ethereum")
	generateAndPrintRandomAddress("Cardano", "cardano")
	generateAndPrintRandomAddress("Litecoin", "litecoin")
	generateAndPrintRandomAddress("Dogecoin", "dogecoin")
	generateAndPrintRandomAddress("Bitcoin Cash", "bitcoincash")
	generateAndPrintRandomAddress("Tron", "tron")
	generateAndPrintRandomAddress("Polkadot", "polkadot")
	generateAndPrintRandomAddress("Ripple", "ripple")

	fmt.Println("\n✅ Done!")
}

// Helper function to generate and print HD-derived address
func generateAndPrintHDAddress(
	wallet *util.CryptoWallet,
	coinName string,
	generator func(uint32) (*util.CryptoAddress, error),
	index uint32,
) {
	addr, err := generator(index)
	if err != nil {
		log.Printf("⚠️ Failed to generate %s HD address: %v", coinName, err)
		return
	}
	fmt.Printf("\n%s (HD Derived) #%d:\n", coinName, index)
	fmt.Printf("🔑 Private Key: %s\n", addr.PrivateKeyHex)
	fmt.Printf("🌐 Public Key:  %s\n", addr.PublicKeyHex)
	fmt.Printf("📬 Address:     %s\n", addr.Address)
}

// Helper function to generate and print random address
func generateAndPrintRandomAddress(coinName, cryptoType string) {
	addr, err := util.GenerateRandomCryptoAddress(cryptoType)
	if err != nil {
		log.Printf("⚠️ Failed to generate random %s address: %v", coinName, err)
		return
	}
	fmt.Printf("\n%s (Random):\n", coinName)
	fmt.Printf("🔑 Private Key: %s\n", addr.PrivateKeyHex)
	fmt.Printf("🌐 Public Key:  %s\n", addr.PublicKeyHex)
	fmt.Printf("📬 Address:     %s\n", addr.Address)
}
