package commands

import (
	"fmt"

	"github.com/aliftech/locksmith/internal/core/lib"
	"github.com/spf13/cobra"
)

var VerCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the current CLI app version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(lib.Cyan("Version: %s"), lib.Cyan(lib.Bold("1.0.0")))
	},
}
