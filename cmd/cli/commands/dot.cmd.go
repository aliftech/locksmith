package commands

import (
	"fmt"

	"github.com/aliftech/locksmith/internal/core/util"
	"github.com/spf13/cobra"
)

var GeneratePolkadotWallet = &cobra.Command{
	Use:   "dot",
	Short: "Generate Polkadot(DOT) wallet address",
	Long:  "Generate a Polkadot wallet address standard with passphrase",
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

		dotWallet, dotErr := wallet.GeneratePolkadotAddress(index)
		if dotErr != nil {
			fmt.Println(red(fmt.Sprintf("ERROR: %s", err)))
			return
		}

		fmt.Println(cyan("Polkadot(DOT) Wallet Address:"))
		fmt.Println(cyan("Mnemonic: ", wallet.Mnemonic))
		fmt.Println(cyan("Public Key: ", dotWallet.PublicKeyHex))
		fmt.Println(cyan("Private Key: ", dotWallet.PrivateKeyHex))
		fmt.Println(cyan("Wallet Address: ", dotWallet.Address))
	},
}

func init() {
	GeneratePolkadotWallet.Flags().StringP("passphrase", "p", "", "Passphrase for generate Polkadot wallet address(required)")
	GeneratePolkadotWallet.Flags().Uint32P("index", "i", 0, "Index for deriving wallet address")
}
