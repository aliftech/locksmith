package commands

import (
	"fmt"

	"github.com/aliftech/locksmith/internal/core/lib"
	"github.com/aliftech/locksmith/internal/core/service"
	"github.com/aliftech/locksmith/internal/core/util"
	"github.com/spf13/cobra"
)

var GenerateBTCPairKeys = &cobra.Command{
	Use:   "btckey",
	Short: "Generate BTC public and private key",
	Run: func(cmd *cobra.Command, args []string) {
		pairkeys, _ := util.GenerateBTCKey(true)
		fmt.Printf(lib.Cyan("BTC public key: %s \n"), lib.Cyan(lib.Bold(pairkeys.PublicKeyHex)))
		fmt.Printf(lib.Cyan("BTC private key: %s \n"), lib.Cyan(lib.Bold(pairkeys.PrivateKeyHex)))
		fmt.Printf(lib.Cyan("BTC WIF: %s \n"), lib.Cyan(lib.Bold(pairkeys.WIF)))
		fmt.Printf(lib.Cyan("BTC P2PKHAddress: %s \n"), lib.Cyan(lib.Bold(pairkeys.P2PKHAddress)))
		fmt.Printf(lib.Cyan("BTC P2TRAddress: %s \n"), lib.Cyan(lib.Bold(pairkeys.P2TRAddress)))
		fmt.Printf(lib.Cyan("BTC P2WPKHAddress: %s \n"), lib.Cyan(lib.Bold(pairkeys.P2WPKHAddress)))
	},
}

var GenerateBTCWallet = &cobra.Command{
	Use:   "btc",
	Short: "Generate Bitcoin(BTC) wallet address",
	Long:  "Generate a Bitcoin wallet address using BIP-44, BIP-84, or BIP-86 standards with passphrase",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		bip44, _ := cmd.Flags().GetBool("bip44")
		bip84, _ := cmd.Flags().GetBool("bip84")
		bip86, _ := cmd.Flags().GetBool("bip86")
		passphrase, _ := cmd.Flags().GetString("passphrase")
		index, _ := cmd.Flags().GetUint32("index")
		saveRemote, _ := cmd.Flags().GetBool("save-remote")

		var bipType string
		var purpose uint32

		// Determine which BIP standard to use
		switch {
		case bip44:
			bipType = "44"
			purpose = 0x8000002C
		case bip84:
			bipType = "84"
			purpose = 0x80000054
		case bip86:
			bipType = "86"
			purpose = 0x80000056
		default:
			fmt.Println(lib.Red("ERROR: Please specify one of --bip44, --bip84, or --bip86"))
			return
		}

		if passphrase == "" {
			fmt.Println(lib.Red("ERROR: passphrase required!"))
			return
		}

		// Generate wallet with the specified BIP standard and passphrase
		wallet, err := util.NewWallet(passphrase)
		if err != nil {
			fmt.Println(lib.Red(fmt.Sprintf("ERROR: %s", err)))
			return
		}

		// Generate BTC wallet address
		walletAddr, err := wallet.GenerateBTCWalletKey(purpose, index)
		if err != nil {
			fmt.Println(lib.Red(fmt.Sprintf("ERROR: %s", err)))
			return
		}

		fmt.Println(lib.Cyan("Bitcoin(BTC) Wallet Address:"))
		fmt.Println(lib.Cyan(fmt.Sprintf("Mnemonic (BIP-%s): %s", bipType, wallet.Mnemonic)))
		switch purpose {
		case 0x8000002C:
			fmt.Println(lib.Cyan(fmt.Sprintf("Public Key: %s", walletAddr.PublicKeyHex)))
			fmt.Println(lib.Cyan(fmt.Sprintf("Private Key: %s", walletAddr.PrivateKeyHex)))
			fmt.Println(lib.Cyan(fmt.Sprintf("Wallet Address (BIP-%s): %s", bipType, walletAddr.P2PKHAddress)))
		case 0x80000054:
			fmt.Println(lib.Cyan(fmt.Sprintf("Public Key: %s", walletAddr.PublicKeyHex)))
			fmt.Println(lib.Cyan(fmt.Sprintf("Private Key: %s", walletAddr.PrivateKeyHex)))
			fmt.Println(lib.Cyan(fmt.Sprintf("Wallet  Address (BIP-%s): %s", bipType, walletAddr.P2WPKHAddress)))
		case 0x80000056:
			fmt.Println(lib.Cyan(fmt.Sprintf("Public Key: %s", walletAddr.PublicKeyHex)))
			fmt.Println(lib.Cyan(fmt.Sprintf("Private Key: %s", walletAddr.PrivateKeyHex)))
			fmt.Println(lib.Cyan(fmt.Sprintf("Wallet  Address (BIP-%s): %s", bipType, walletAddr.P2TRAddress)))
		default:
			fmt.Println(lib.Cyan(fmt.Sprintf("Public Key: %s", walletAddr.PublicKeyHex)))
			fmt.Println(lib.Cyan(fmt.Sprintf("Private Key: %s", walletAddr.PrivateKeyHex)))
			fmt.Println(lib.Cyan(fmt.Sprintf("Wallet Address: %s", walletAddr.WIF)))
		}

		// Implement gRPC
		if saveRemote {
			if err := service.StoreBTCWalletViaGRPC(wallet.Mnemonic, "BTC", walletAddr, index, passphrase, purpose); err != nil {
				fmt.Println(lib.Red(fmt.Sprintf("gRPC Save ERROR: %s", err)))
			} else {
				fmt.Println(lib.Green(lib.Bold("✅ Wallet saved remotely via gRPC")))
			}
		}
	},
}

func init() {
	GenerateBTCWallet.Flags().Bool("bip44", false, "Generate wallet using BIP-44 (P2PKH)")
	GenerateBTCWallet.Flags().Bool("bip84", false, "Generate wallet using BIP-84 (P2WPKH)")
	GenerateBTCWallet.Flags().Bool("bip86", false, "Generate wallet using BIP-86 (P2TR)")
	GenerateBTCWallet.Flags().StringP("passphrase", "p", "", "Passphrase for wallet generation (required)")
	GenerateBTCWallet.Flags().Uint32P("index", "i", 0, "Index for deriving wallet address (default 0)")
	GenerateBTCWallet.Flags().Bool("save-remote", false, "Save wallet to remote gRPC server")
}
