package cmd

import (
	"fmt"
	"os"

	"github.com/c88lopez/dbs/src/config"

	"github.com/c88lopez/dbs/src/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:  "dbs",
	Long: "A database schema changes",

	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
	SilenceErrors: true,
	SilenceUsage:  true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		errors.Handle(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find main folder directory.
		mainFolder, err := config.GetMainFolderPath()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name "config" (without extension).
		viper.AddConfigPath(mainFolder)
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")

		viper.SetDefault("Driver", "mysql")
		viper.SetDefault("Protocol", "tcp")
		viper.SetDefault("Host", "127.0.0.1")
		viper.SetDefault("Port", "3306")
		viper.SetDefault("Username", "root")
		viper.SetDefault("Password", "")
		viper.SetDefault("Database", "dbs")

		err = viper.SafeWriteConfig()
		if err != nil {
			errors.Handle(err)

		}
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
