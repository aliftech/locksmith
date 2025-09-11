package main

import (
	"fmt"

	"github.com/aliftech/locksmith/internal/core/commands"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var border = color.New(color.FgCyan).SprintFunc()
var title = color.New(color.FgCyan, color.Bold).SprintFunc()

var cmd = &cobra.Command{
	Use: title("locksmith"),
	Annotations: map[string]string{
		cobra.CommandDisplayNameAnnotation: title("locksmith"),
	},
}

func main() {
	fmt.Println(border("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓"))
	fmt.Printf("┃%s┃\n", title("                  LOCKSMITH                   "))
	fmt.Println(border("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛"))

	cmd.AddCommand(commands.GenerateBTCWallet)
	cmd.AddCommand(commands.GenerateBitcoinCashWallet)
	cmd.AddCommand(commands.GenerateEthWallet)
	cmd.AddCommand(commands.GenerateCardanoWallet)
	cmd.AddCommand(commands.GenerateLitecoinWallet)
	cmd.AddCommand(commands.GenerateDogecoinWallet)
	cmd.AddCommand(commands.GenerateTronWallet)
	cmd.AddCommand(commands.GeneratePolkadotWallet)
	cmd.AddCommand(commands.GenerateRippleWallet)
	cmd.AddCommand(commands.VerCmd)
	if err := cmd.Execute(); err != nil {
		color.RedString("ERROR: %s", err)
	}
}
