package commands

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"

	"github.com/aliftech/locksmith/internal/core/lib"
	"github.com/aliftech/locksmith/internal/core/service"
	"github.com/aliftech/locksmith/internal/core/util"
)

var GenerateEthWallet = &cobra.Command{
	Use:   "eth",
	Short: "Generate Etherium(ETH) wallet address",
	Long:  "Generate an Etherium wallet address standard with passphrase",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		passphrase, _ := cmd.Flags().GetString("passphrase")
		index, _ := cmd.Flags().GetUint32("index")
		saveRemote, _ := cmd.Flags().GetBool("save-remote") // Optional flag

		if passphrase == "" {
			fmt.Println(lib.Red("ERROR: passphrase required!"))
			return
		}

		walletAddr, err := util.NewCryptoWallet(passphrase)
		if err != nil {
			fmt.Println(lib.Red(fmt.Sprintf("ERROR: %s", err)))
			return
		}

		ethWallet, ethErr := walletAddr.GenerateEthereumAddress(index)
		if ethErr != nil {
			fmt.Println(lib.Red(fmt.Sprintf("ERROR: %s", err)))
			return
		}

		fmt.Println(lib.Cyan("Etherium(ETH) Wallet Address:"))
		fmt.Println(lib.Cyan("Mnemonic: ", walletAddr.Mnemonic))
		fmt.Println(lib.Cyan("Public Key: ", ethWallet.PublicKeyHex))
		fmt.Println(lib.Cyan("Private Key: ", ethWallet.PrivateKeyHex))
		fmt.Println(lib.Cyan("Wallet Address: ", ethWallet.Address))

		// Optionally store via gRPC
		if saveRemote {
			if err := service.StoreWalletViaGRPC(walletAddr.Mnemonic, "eth", ethWallet, index, passphrase); err != nil {
				fmt.Println(lib.Red(fmt.Sprintf("gRPC Save ERROR: %s", err)))
			} else {
				fmt.Println(lib.Green(lib.Bold("✅ Wallet saved remotely via gRPC")))
			}
		}
	},
}

func init() {
	GenerateEthWallet.Flags().StringP("passphrase", "p", "", "Passphrase for generate Etherium wallet address(required)")
	GenerateEthWallet.Flags().Uint32P("index", "i", 0, "Index for deriving wallet address")
	GenerateEthWallet.Flags().Bool("save-remote", false, "Save wallet to remote gRPC server")
}
