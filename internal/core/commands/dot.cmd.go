package commands

import (
	"fmt"

	"github.com/aliftech/locksmith/internal/core/lib"
	"github.com/aliftech/locksmith/internal/core/service"
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

		dotWallet, dotErr := wallet.GeneratePolkadotAddress(index)
		if dotErr != nil {
			fmt.Println(lib.Red(fmt.Sprintf("ERROR: %s", err)))
			return
		}

		fmt.Println(lib.Cyan("Polkadot(DOT) Wallet Address:"))
		fmt.Println(lib.Cyan("Mnemonic: ", wallet.Mnemonic))
		fmt.Println(lib.Cyan("Public Key: ", dotWallet.PublicKeyHex))
		fmt.Println(lib.Cyan("Private Key: ", dotWallet.PrivateKeyHex))
		fmt.Println(lib.Cyan("Wallet Address: ", dotWallet.Address))

		// Implement gRPC server
		if saveRemote {
			if err := service.StoreWalletViaGRPC(wallet.Mnemonic, "DOT", dotWallet, index, passphrase); err != nil {
				fmt.Println(lib.Red(fmt.Sprintf("gRPC Save ERROR: %s", err)))
			} else {
				fmt.Println(lib.Green(lib.Bold("✅ Wallet saved remotely via gRPC")))
			}
		}
	},
}

func init() {
	GeneratePolkadotWallet.Flags().StringP("passphrase", "p", "", "Passphrase for generate Polkadot wallet address(required)")
	GeneratePolkadotWallet.Flags().Uint32P("index", "i", 0, "Index for deriving wallet address")
	GeneratePolkadotWallet.Flags().Bool("save-remote", false, "Save wallet to remote gRPC server")
}
