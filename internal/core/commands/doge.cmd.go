package commands

import (
	"fmt"

	"github.com/aliftech/locksmith/internal/core/lib"
	"github.com/aliftech/locksmith/internal/core/service"
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

		dogeWallet, dogeErr := wallet.GenerateDogecoinAddress(index)
		if dogeErr != nil {
			fmt.Println(lib.Red(fmt.Sprintf("ERROR: %s", err)))
			return
		}

		fmt.Println(lib.Cyan("Dogecoin(DOGE) Wallet Address:"))
		fmt.Println(lib.Cyan("Mnemonic: ", wallet.Mnemonic))
		fmt.Println(lib.Cyan("Public Key: ", dogeWallet.PublicKeyHex))
		fmt.Println(lib.Cyan("Private Key: ", dogeWallet.PrivateKeyHex))
		fmt.Println(lib.Cyan("Wallet Address: ", dogeWallet.Address))

		// Implement gRPC
		if saveRemote {
			if err := service.StoreWalletViaGRPC(wallet.Mnemonic, "DOGE", dogeWallet, index, passphrase); err != nil {
				fmt.Println(lib.Red(fmt.Sprintf("gRPC Save ERROR: %s", err)))
			} else {
				fmt.Println(lib.Green(lib.Bold("✅ Wallet saved remotely via gRPC")))
			}
		}
	},
}

func init() {
	GenerateDogecoinWallet.Flags().StringP("passphrase", "p", "", "Passphrase for generate Dogecoin standard address(required)")
	GenerateDogecoinWallet.Flags().Uint32P("index", "i", 0, "Index for deriving wallet address")
	GenerateDogecoinWallet.Flags().Bool("save-remote", false, "Save wallet to remote gRPC  server")
}
