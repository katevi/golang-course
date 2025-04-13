package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "A Viper-powered app",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd:   true,
		DisableDescriptions: true,
		DisableNoDescFlag:   true,
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		initConfig() // Initialize Viper before any command runs
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
