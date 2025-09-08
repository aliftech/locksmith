package main

import (
	"fmt"
	"log"

	"github.com/aliftech/locksmith/internal/core/util"
)

func main() {
	// Create a new wallet
	wallet, err := util.NewWallet("optional-passphrase")
	if err != nil {
		log.Fatal(err)
	}

	// Print mnemonic
	fmt.Println("Mnemonic:", wallet.Mnemonic)

	// Generate BIP-44 key (P2PKH)
	bip44Key, err := wallet.GenerateBTCWalletKey(0x8000002C, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("BIP-44 P2PKH Address: %s\n", bip44Key.P2PKHAddress)

	// Generate BIP-84 key (P2WPKH)
	bip84Key, err := wallet.GenerateBTCWalletKey(0x80000054, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("BIP-84 P2WPKH Address: %s\n", bip84Key.P2WPKHAddress)

	// Generate BIP-86 key (P2TR)
	bip86Key, err := wallet.GenerateBTCWalletKey(0x80000056, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("BIP-86 P2TR Address: %s\n", bip86Key.P2TRAddress)

	// Generate random key (non-mnemonic)
	randomKey, err := util.GenerateBTCKey(true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Random P2PKH Address: %s\n", randomKey.P2PKHAddress)
}
