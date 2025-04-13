package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var morningCmd = &cobra.Command{
	Use:   "morning",
	Short: "Morning greeting",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if verbose {
			fmt.Printf("Sending greeting to: %s\n", name)
		}
		fmt.Printf("Hello, %s!\n", name)
	},
}

func init() {
	morningCmd.Flags().StringP("name", "n", "World", "Name to greet")
	greetCmd.AddCommand(morningCmd)
}
