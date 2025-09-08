package main

import (
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use: "locksmith",
		Annotations: map[string]string{
			cobra.CommandDisplayNameAnnotation: "locksmith",
		},
	}
	cmd.Execute()
}
