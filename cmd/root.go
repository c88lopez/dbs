package cmd

import (
	"fmt"

	"github.com/c88lopez/dbs/src/errors"
	"github.com/c88lopez/dbs/src/mainFolder"
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
	errors.Handle(rootCmd.Execute())
}

func init() {
	// 	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file if set.
func initConfig() {
	viper.AddConfigPath(mainFolder.GetMainFolderPath())
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.SafeWriteConfig()
	if err != nil {
		errors.Handle(err)
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
