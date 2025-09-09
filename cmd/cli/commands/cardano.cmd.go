package commands

import (
	"fmt"

	"github.com/aliftech/locksmith/internal/core/util"
	"github.com/spf13/cobra"
)

var GenerateCardanoWallet = &cobra.Command{
	Use:   "ada",
	Short: "Generate Cardano(ADA) wallet address",
	Long:  "Generate a cardano(ADA) wallet address standard with passphrase",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		passphrase, _ := cmd.Flags().GetString("passphrase")
		index, _ := cmd.Flags().GetUint32("index")

		if passphrase == "" {
			fmt.Println(red("ERROR: passphrase required!"))
			return
		}

		walletAddr, err := util.NewCryptoWallet(passphrase)
		if err != nil {
			fmt.Println(red(fmt.Sprintf("ERROR: %s", err)))
			return
		}

		adaWallet, adaErr := walletAddr.GenerateCardanoAddress(index)
		if adaErr != nil {
			fmt.Println(red(fmt.Sprintf("ERROR: %s", err)))
			return
		}

		fmt.Println(cyan("Cardano(ADA) Wallet Address:"))
		fmt.Println(cyan("Mnemonic: ", walletAddr.Mnemonic))
		fmt.Println(cyan("Public Key: ", adaWallet.PublicKeyHex))
		fmt.Println(cyan("Private Key: ", adaWallet.PrivateKeyHex))
		fmt.Println(cyan("Wallet Address: ", adaWallet.Address))
	},
}

func init() {
	GenerateCardanoWallet.Flags().StringP("passphrase", "p", "", "Passphrase for generate cardano wallet address(required)")
	GenerateCardanoWallet.Flags().Uint32P("index", "i", 0, "Index for deriving wallet address")
}
