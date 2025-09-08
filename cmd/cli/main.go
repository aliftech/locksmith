package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/aliftech/locksmith/cmd/cli/commands"
)

var cmd = &cobra.Command{
	Use: "locksmith",
	Annotations: map[string]string{
		cobra.CommandDisplayNameAnnotation: "locksmith",
	},
}

func main() {
	border := color.New(color.FgCyan).SprintFunc()
	title := color.New(color.FgCyan, color.Bold).SprintFunc()

	fmt.Println(border("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓"))
	fmt.Printf("┃%s┃\n", title("                  LOCKSMITH                   "))
	fmt.Println(border("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛"))

	cmd.AddCommand(commands.VerCmd)
	if err := cmd.Execute(); err != nil {
		color.RedString("ERROR (%s)", err)
	}
}
