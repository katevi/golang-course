package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Greet someone",
	Long:  `This command greets a person by name`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		fmt.Printf("Hello, %s!\n", name)
	},
}

func init() {
	greetCmd.Flags().StringP("name", "n", "World", "Name to greet")
	rootCmd.AddCommand(greetCmd)
}
