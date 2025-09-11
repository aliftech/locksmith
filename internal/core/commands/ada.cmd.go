package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/aliftech/locksmith/internal/core/lib"
	"github.com/aliftech/locksmith/internal/core/util"
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
			fmt.Println(lib.Red("ERROR: passphrase required!"))
			return
		}

		walletAddr, err := util.NewCryptoWallet(passphrase)
		if err != nil {
			fmt.Println(lib.Red(fmt.Sprintf("ERROR: %s", err)))
			return
		}

		adaWallet, adaErr := walletAddr.GenerateCardanoAddress(index)
		if adaErr != nil {
			fmt.Println(lib.Red(fmt.Sprintf("ERROR: %s", err)))
			return
		}

		fmt.Println(lib.Red(lib.Bold("Cardano(ADA) Wallet Address:")))
		fmt.Println(lib.Cyan("Mnemonic: ", walletAddr.Mnemonic))
		fmt.Println(lib.Cyan("Public Key: ", adaWallet.PublicKeyHex))
		fmt.Println(lib.Cyan("Private Key: ", adaWallet.PrivateKeyHex))
		fmt.Println(lib.Cyan("Wallet Address: ", adaWallet.Address))
	},
}

func init() {
	GenerateCardanoWallet.Flags().StringP("passphrase", "p", "", "Passphrase for generate cardano wallet address(required)")
	GenerateCardanoWallet.Flags().Uint32P("index", "i", 0, "Index for deriving wallet address")
}
