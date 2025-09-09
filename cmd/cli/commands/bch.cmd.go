package commands

import (
	"fmt"

	"github.com/aliftech/locksmith/internal/core/util"
	"github.com/spf13/cobra"
)

var GenerateBitcoinCashWallet = &cobra.Command{
	Use:   "bch",
	Short: "Generate Bitcoin Cash(BCH) wallet address",
	Long:  "Generate a Bitcoin Cash(BCH) wallet address standard with passphrase",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		passphrase, _ := cmd.Flags().GetString("passphrase")
		index, _ := cmd.Flags().GetUint32("index")

		if passphrase == "" {
			fmt.Println(red("ERROR: passphrase required!"))
			return
		}

		wallet, err := util.NewCryptoWallet(passphrase)
		if err != nil {
			fmt.Println(red(fmt.Sprintf("ERROR: %s", err)))
			return
		}

		bchWallet, bchErr := wallet.GenerateBitcoinCashAddress(index)
		if bchErr != nil {
			fmt.Println(red(fmt.Sprintf("ERROR: %s", err)))
			return
		}

		fmt.Println(cyan("Bitcoin Cash(BCH) Wallet Address:"))
		fmt.Println(cyan("Mnemonic: ", wallet.Mnemonic))
		fmt.Println(cyan("Public Key: ", bchWallet.PublicKeyHex))
		fmt.Println(cyan("Private Key: ", bchWallet.PrivateKeyHex))
		fmt.Println(cyan("Wallet Address: ", bchWallet.Address))
	},
}

func init() {
	GenerateBitcoinCashWallet.Flags().StringP("passphrase", "p", "", "Passphrase for generate Bitcoin Cash wallet address(required)")
	GenerateBitcoinCashWallet.Flags().Uint32P("index", "i", 0, "Index for deriving wallet address")
}
