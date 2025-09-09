package commands

import (
	"context"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/aliftech/locksmith/internal/core/util"
	walletpb "github.com/aliftech/locksmith/internal/grpc"
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

		// Optionally store via gRPC
		if saveRemote {
			if err := storeWalletViaGRPC(walletAddr.Mnemonic, ethWallet, index, passphrase); err != nil {
				fmt.Println(red(fmt.Sprintf("gRPC Save ERROR: %s", err)))
			} else {
				fmt.Println(cyan("✅ Wallet saved remotely via gRPC"))
			}
		}
	},
}

func init() {
	GenerateEthWallet.Flags().StringP("passphrase", "p", "", "Passphrase for generate Etherium wallet address(required)")
	GenerateEthWallet.Flags().Uint32P("index", "i", 0, "Index for deriving wallet address")
	GenerateEthWallet.Flags().Bool("save-remote", false, "Save wallet to remote gRPC server")
}

func storeWalletViaGRPC(mnemonic string, ethWallet *util.CryptoAddress, index uint32, passphrase string) error {
	// Connect to gRPC server
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}
	defer conn.Close()

	client := walletpb.NewWalletServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Send request
	req := &walletpb.StoreWalletRequest{
		Mnemonic:       mnemonic,
		PublicKey:      ethWallet.PublicKeyHex,
		PrivateKey:     ethWallet.PrivateKeyHex,
		Address:        ethWallet.Address,
		Index:          index,
		PassphraseHash: passphrase, // Server will hash it
	}

	res, err := client.StoreWallet(ctx, req)
	if err != nil {
		return fmt.Errorf("RPC failed: %w", err)
	}

	if !res.Success {
		return fmt.Errorf("server error: %s", res.Message)
	}

	return nil
}
