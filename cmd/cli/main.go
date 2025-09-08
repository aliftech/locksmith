package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/aliftech/locksmith/cmd/cli/commands"
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

	cmd.AddCommand(commands.GenerateBTCPairKeys)
	cmd.AddCommand(commands.GenerateBTCWallet)
	cmd.AddCommand(commands.VerCmd)
	if err := cmd.Execute(); err != nil {
		color.RedString("ERROR: %s", err)
	}
}
