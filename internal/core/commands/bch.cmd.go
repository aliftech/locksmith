package commands

import (
	"fmt"

	"github.com/aliftech/locksmith/internal/core/app/services"
	"github.com/aliftech/locksmith/internal/core/lib"
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
		saveRemote, _ := cmd.Flags().GetBool("save-remote")

		if passphrase == "" {
			fmt.Println(lib.Red("ERROR: passphrase required!"))
			return
		}

		wallet, err := util.NewCryptoWallet(passphrase)
		if err != nil {
			fmt.Println(lib.Red(fmt.Sprintf("ERROR: %s", err)))
			return
		}

		bchWallet, bchErr := wallet.GenerateBitcoinCashAddress(index)
		if bchErr != nil {
			fmt.Println(lib.Red(fmt.Sprintf("ERROR: %s", err)))
			return
		}

		fmt.Println(lib.Cyan("Bitcoin Cash(BCH) Wallet Address:"))
		fmt.Println(lib.Cyan("Mnemonic: ", wallet.Mnemonic))
		fmt.Println(lib.Cyan("Public Key: ", bchWallet.PublicKeyHex))
		fmt.Println(lib.Cyan("Private Key: ", bchWallet.PrivateKeyHex))
		fmt.Println(lib.Cyan("Wallet Address: ", bchWallet.Address))

		if saveRemote {
			if err := services.StoreWalletViaGRPC(wallet.Mnemonic, "BCH", bchWallet, index, passphrase); err != nil {
				fmt.Println(lib.Red(fmt.Sprintf("gRPC Save ERROR: %s", err)))
			} else {
				fmt.Println(lib.Green(lib.Bold("✅ Wallet saved remotely via gRPC")))
			}
		}
	},
}

func init() {
	GenerateBitcoinCashWallet.Flags().StringP("passphrase", "p", "", "Passphrase for generate Bitcoin Cash wallet address(required)")
	GenerateBitcoinCashWallet.Flags().Uint32P("index", "i", 0, "Index for deriving wallet address")
	GenerateBitcoinCashWallet.Flags().Bool("save-remote", false, "Save wallet to remote gRPC")
}
