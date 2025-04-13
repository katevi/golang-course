package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var verbose bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}

var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "A simple CLI application",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd:   true,
		DisableDescriptions: true,
		DisableNoDescFlag:   true,
	},
	Long: `A longer description of my CLI application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to my app! Use --help to see commands")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
