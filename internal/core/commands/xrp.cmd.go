package commands

import (
	"fmt"

	"github.com/aliftech/locksmith/internal/core/lib"
	"github.com/aliftech/locksmith/internal/core/service"
	"github.com/aliftech/locksmith/internal/core/util"
	"github.com/spf13/cobra"
)

var GenerateRippleWallet = &cobra.Command{
	Use:   "xrp",
	Short: "Generate Ripple(XRP) wallet address",
	Long:  "Generate a Ripple wallet address standard with passphrase",
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

		xrpWallet, xrpErr := wallet.GenerateRippleAddress(index)
		if xrpErr != nil {
			fmt.Println(lib.Red(fmt.Sprintf("ERROR: %s", err)))
			return
		}

		fmt.Println(lib.Cyan("Ripple(XRP) Wallet Address:"))
		fmt.Println(lib.Cyan("Mnemonic: ", wallet.Mnemonic))
		fmt.Println(lib.Cyan("Public Key: ", xrpWallet.PublicKeyHex))
		fmt.Println(lib.Cyan("Private Key: ", xrpWallet.PrivateKeyHex))
		fmt.Println(lib.Cyan("Wallet Address: ", xrpWallet.Address))

		// Implement gRPC server
		if saveRemote {
			if err := service.StoreWalletViaGRPC(wallet.Mnemonic, "XRP", xrpWallet, index, passphrase); err != nil {
				fmt.Println(lib.Red(fmt.Sprintf("gRPC Save ERROR: %s", err)))
			} else {
				fmt.Println(lib.Green(lib.Bold("✅ Wallet saved remotely via gRPC")))
			}
		}
	},
}

func init() {
	GenerateRippleWallet.Flags().StringP("passphrase", "p", "", "Passphrase for generate Rippler wallet address(required)")
	GenerateRippleWallet.Flags().Uint32P("index", "i", 0, "Index for deriving wallet address")
	GenerateRippleWallet.Flags().Bool("save-remote", false, "Save wallet to remote gRPC server")
}
