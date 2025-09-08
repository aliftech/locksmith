package commands

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var cyan = color.New(color.FgCyan).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var title = color.New(color.FgCyan, color.Bold).SprintFunc()

var VerCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the current CLI app version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(title("Version: %s"), cyan("1.0.0"))
	},
}
