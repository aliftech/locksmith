package commands

import (
	"fmt"

	"github.com/aliftech/locksmith/internal/core/util"
	"github.com/spf13/cobra"
)

var GenerateEthWallet = &cobra.Command{
	Use:   "eth",
	Short: "Generate Etherium(ETH) wallet address",
	Long:  "Generate an Etherium wallet address standard with passphrase",
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

		ethWallet, ethErr := walletAddr.GenerateEthereumAddress(index)
		if ethErr != nil {
			fmt.Println(red(fmt.Sprintf("ERROR: %s", err)))
			return
		}

		fmt.Println(cyan("Etherium(ETH) Wallet Address:"))
		fmt.Println(cyan("Mnemonic: ", walletAddr.Mnemonic))
		fmt.Println(cyan("Public Key: ", ethWallet.PublicKeyHex))
		fmt.Println(cyan("Private Key: ", ethWallet.PrivateKeyHex))
		fmt.Println(cyan("Wallet Address: ", ethWallet.Address))
	},
}

func init() {
	GenerateEthWallet.Flags().StringP("passphrase", "p", "", "Passphrase for generate Etherium wallet address(required)")
	GenerateEthWallet.Flags().Uint32P("index", "i", 0, "Index for deriving wallet address")
}
