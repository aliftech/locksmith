package commands

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var shortDescription = color.New(color.FgCyan).SprintFunc()
var verTitle = color.New(color.FgCyan, color.Bold).SprintFunc()

var VerCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the current CLI app version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(verTitle("Version: %s"), shortDescription("1.0.0"))
	},
}
