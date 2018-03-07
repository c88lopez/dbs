package cmd

import (
	"fmt"

	"github.com/c88lopez/dbs/src/errors"
	"github.com/c88lopez/dbs/src/mainfolder"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the dbs.",
	Long:  "Configure the dbs.",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		initConfig()

		err := viper.ReadInConfig()
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err = errors.ErrNotInitialized
		}

		return err
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print("Configuring database parameters...\n")

		return mainfolder.SetDatabaseConfigInteractively()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
