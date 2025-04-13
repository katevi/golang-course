package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	serverCmd.Flags().StringP("port", "p", "", "Server port")
	viper.BindPFlag("server.port", serverCmd.Flags().Lookup("port"))

	// Command-line will override config file:
	// ./myapp start -p 4000
}

// Example command using Viper values
var serverCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		port := viper.GetString("server.port")
		logLevel := viper.GetString("logging.level")

		fmt.Printf("Starting server on %s with log level %s\n",
			port, logLevel)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
