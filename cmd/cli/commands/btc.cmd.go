package commands

import (
	"fmt"

	"github.com/aliftech/locksmith/internal/core/util"
	"github.com/spf13/cobra"
)

var GenerateBTCPairKeys = &cobra.Command{
	Use:   "btckey",
	Short: "Generate BTC public and private key",
	Run: func(cmd *cobra.Command, args []string) {
		pairkeys, _ := util.GenerateBTCKey(true)
		fmt.Printf(cyan("BTC public key: %s \n"), title(pairkeys.PublicKeyHex))
		fmt.Printf(cyan("BTC private key: %s \n"), title(pairkeys.PrivateKeyHex))
		fmt.Printf(cyan("BTC WIF: %s \n"), title(pairkeys.WIF))
		fmt.Printf(cyan("BTC P2PKHAddress: %s \n"), title(pairkeys.P2PKHAddress))
		fmt.Printf(cyan("BTC P2TRAddress: %s \n"), title(pairkeys.P2TRAddress))
		fmt.Printf(cyan("BTC P2WPKHAddress: %s \n"), title(pairkeys.P2WPKHAddress))
	},
}

var GenerateBTCWallet = &cobra.Command{
	Use:   "btc",
	Short: "Generate BTC Wallet Address",
	Long:  "Generate a BTC wallet address using BIP-44, BIP-84, or BIP-86 standards with an optional passphrase",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		bip44, _ := cmd.Flags().GetBool("bip44")
		bip84, _ := cmd.Flags().GetBool("bip84")
		bip86, _ := cmd.Flags().GetBool("bip86")
		passphrase, _ := cmd.Flags().GetString("passphrase")
		index, _ := cmd.Flags().GetUint32("index")

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
			fmt.Println(red("ERROR: Please specify one of --bip44, --bip84, or --bip86"))
			return
		}

		if passphrase == "" {
			fmt.Println(red("ERROR: passphrase required!"))
			return
		}

		// Generate wallet with the specified BIP standard and passphrase
		wallet, err := util.NewWallet(passphrase)
		if err != nil {
			fmt.Println(red(fmt.Sprintf("ERROR: %s", err)))
			return
		}

		// Generate BTC wallet address
		walletAddr, err := wallet.GenerateBTCWalletKey(purpose, index)
		if err != nil {
			fmt.Println(red(fmt.Sprintf("ERROR: %s", err)))
			return
		}

		fmt.Println(cyan(fmt.Sprintf("Your BTC mnemonic (BIP-%s): %s", bipType, wallet.Mnemonic)))
		switch purpose {
		case 0x8000002C:
			fmt.Println(cyan(fmt.Sprintf("Your BTC public key: %s", walletAddr.PublicKeyHex)))
			fmt.Println(cyan(fmt.Sprintf("Your BTC private key: %s", walletAddr.PrivateKeyHex)))
			fmt.Println(cyan(fmt.Sprintf("Your BTC wallet address (BIP-%s): %s", bipType, walletAddr.P2PKHAddress)))
		case 0x80000054:
			fmt.Println(cyan(fmt.Sprintf("Your BTC public key: %s", walletAddr.PublicKeyHex)))
			fmt.Println(cyan(fmt.Sprintf("Your BTC private key: %s", walletAddr.PrivateKeyHex)))
			fmt.Println(cyan(fmt.Sprintf("Your BTC wallet address (BIP-%s): %s", bipType, walletAddr.P2WPKHAddress)))
		case 0x80000056:
			fmt.Println(cyan(fmt.Sprintf("Your BTC public key: %s", walletAddr.PublicKeyHex)))
			fmt.Println(cyan(fmt.Sprintf("Your BTC private key: %s", walletAddr.PrivateKeyHex)))
			fmt.Println(cyan(fmt.Sprintf("Your BTC wallet address (BIP-%s): %s", bipType, walletAddr.P2TRAddress)))
		default:
			fmt.Println(cyan(fmt.Sprintf("Your BTC public key: %s", walletAddr.PublicKeyHex)))
			fmt.Println(cyan(fmt.Sprintf("Your BTC private key: %s", walletAddr.PrivateKeyHex)))
			fmt.Println(cyan(fmt.Sprintf("Your WIF BTC wallet address: %s", walletAddr.WIF)))
		}
	},
}

func init() {
	GenerateBTCWallet.Flags().Bool("bip44", false, "Generate wallet using BIP-44 (P2PKH)")
	GenerateBTCWallet.Flags().Bool("bip84", false, "Generate wallet using BIP-84 (P2WPKH)")
	GenerateBTCWallet.Flags().Bool("bip86", false, "Generate wallet using BIP-86 (P2TR)")
	GenerateBTCWallet.Flags().StringP("passphrase", "p", "", "Passphrase for wallet generation (required)")
	GenerateBTCWallet.Flags().Uint32P("index", "i", 0, "Index for deriving wallet address (default 0)")
}
