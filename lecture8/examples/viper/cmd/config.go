package cmd

import (
	"fmt"

	"github.com/spf13/viper"
)

func initConfig() {
	// 1. Set config file name and paths
	viper.SetConfigName("config")    // Looks for config.yaml/config.json/etc
	viper.AddConfigPath(".")         // Current directory
	viper.AddConfigPath("./configs") // Configs subfolder

	// 2. Set default values
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("logging.level", "info")

	viper.AutomaticEnv()        // Read ENV variables
	viper.SetEnvPrefix("MYAPP") // MYAPP_SERVER_PORT=5000

	// Bind ENV to config keys:
	viper.BindEnv("server.port")

	// 3. Read config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("No config file found - using defaults")
		} else {
			panic(fmt.Errorf("Fatal error in config file: %w", err))
		}
	}
}
