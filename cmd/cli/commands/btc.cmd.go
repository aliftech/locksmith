package commands

import (
	"fmt"

	"github.com/aliftech/locksmith/internal/core/util"
	"github.com/spf13/cobra"
)

var pairkeys, _ = util.GenerateBTCKey(true)

var GenerateBTCPairKeys = &cobra.Command{
	Use:   "btckey",
	Short: "Generate BTC public and private key",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(cyan("BTC public key: %s \n"), title(pairkeys.PublicKeyHex))
		fmt.Printf(cyan("BTC private key: %s \n"), title(pairkeys.PrivateKeyHex))
		fmt.Printf(cyan("BTC WIF: %s \n"), title(pairkeys.WIF))
		fmt.Printf(cyan("BTC P2PKHAddress: %s \n"), title(pairkeys.P2PKHAddress))
		fmt.Printf(cyan("BTC P2TRAddress: %s \n"), title(pairkeys.P2TRAddress))
		fmt.Printf(cyan("BTC P2WPKHAddress: %s \n"), title(pairkeys.P2WPKHAddress))
	},
}
