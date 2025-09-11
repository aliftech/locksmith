package commands

import (
	"fmt"

	"github.com/aliftech/locksmith/internal/core/lib"
	"github.com/aliftech/locksmith/internal/core/util"
	"github.com/spf13/cobra"
)

var GenerateLitecoinWallet = &cobra.Command{
	Use:   "ltc",
	Short: "Generate Litecoin(LTC) wallet address",
	Long:  "Generate a Litecoin wallet address standard with passphrase",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		passphrase, _ := cmd.Flags().GetString("passphrase")
		index, _ := cmd.Flags().GetUint32("index")

		if passphrase == "" {
			fmt.Println(lib.Red("ERROR: passphrase required!"))
			return
		}

		wallet, err := util.NewCryptoWallet(passphrase)
		if err != nil {
			fmt.Println(lib.Red(fmt.Sprintf("ERROR: %s", err)))
			return
		}

		ltcWallet, ltcErr := wallet.GenerateLitecoinAddress(index)
		if ltcErr != nil {
			fmt.Println(lib.Red(fmt.Sprintf("ERROR: %s", err)))
			return
		}

		fmt.Println(lib.Cyan("Litecoin(LTC) Wallet Address:"))
		fmt.Println(lib.Cyan("Mnemonic: ", wallet.Mnemonic))
		fmt.Println(lib.Cyan("Public Key: ", ltcWallet.PublicKeyHex))
		fmt.Println(lib.Cyan("Private Key: ", ltcWallet.PrivateKeyHex))
		fmt.Println("Wallet Address: ", ltcWallet.Address)
	},
}

func init() {
	GenerateLitecoinWallet.Flags().StringP("passphrase", "p", "", "Passphrase for generate Litecoin wallet address(required)")
	GenerateLitecoinWallet.Flags().Uint32P("index", "i", 0, "Index for deriving wallet address")
}
