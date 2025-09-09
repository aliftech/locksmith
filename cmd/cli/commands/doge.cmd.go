package commands

import (
	"fmt"

	"github.com/aliftech/locksmith/internal/core/util"
	"github.com/spf13/cobra"
)

var GenerateDogecoinWallet = &cobra.Command{
	Use:   "doge",
	Short: "Generate Dogecoin(DOGE) wallet address",
	Long:  "Generate Dogecoin wallet address standard with passphrase",
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

		dogeWallet, dogeErr := wallet.GenerateDogecoinAddress(index)
		if dogeErr != nil {
			fmt.Println(red(fmt.Sprintf("ERROR: %s", err)))
			return
		}

		fmt.Println(cyan("Dogecoin(DOGE) Wallet Address:"))
		fmt.Println(cyan("Mnemonic: ", wallet.Mnemonic))
		fmt.Println(cyan("Public Key: ", dogeWallet.PublicKeyHex))
		fmt.Println(cyan("Private Key: ", dogeWallet.PrivateKeyHex))
		fmt.Println(cyan("Wallet Address: ", dogeWallet.Address))
	},
}

func init() {
	GenerateDogecoinWallet.Flags().StringP("passphrase", "p", "", "Passphrase for generate Dogecoin standard address(required)")
	GenerateDogecoinWallet.Flags().Uint32P("index", "i", 0, "Index for deriving wallet address")
}
