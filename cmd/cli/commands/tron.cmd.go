package commands

import (
	"fmt"

	"github.com/aliftech/locksmith/internal/core/util"
	"github.com/spf13/cobra"
)

var GenerateTronWallet = &cobra.Command{
	Use:   "tron",
	Short: "Generate Tron(TRON) wallet address",
	Long:  "Generate a Tron wallet address standard with passphrase",
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

		tronWallet, tronErr := wallet.GenerateTronAddress(index)
		if tronErr != nil {
			fmt.Println(red(fmt.Sprintf("ERROR: %s", err)))
			return
		}

		fmt.Println(cyan("Tron(TRON) Wallet Address:"))
		fmt.Println(cyan("Mnemonic: ", wallet.Mnemonic))
		fmt.Println(cyan("Public Key: ", tronWallet.PublicKeyHex))
		fmt.Println(cyan("Private Key: ", tronWallet.PrivateKeyHex))
		fmt.Println(cyan("Wallet Address: ", tronWallet.Address))
	},
}

func init() {
	GenerateTronWallet.Flags().StringP("passphrase", "p", "", "Passphrase for generate Tron wallet address(required)")
	GenerateTronWallet.Flags().Uint32P("index", "i", 0, "Index for deriving wallet address")
}
