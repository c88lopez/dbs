package cmd

import (
	"fmt"

	"github.com/c88lopez/dbs/src/mainfolder"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize dbs",
	Long:  `Initialize dbs.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print("Initializing... ")

		return mainfolder.Generate()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	viper.SetDefault("driver", "mysql")
	viper.SetDefault("protocol", "tcp")
	viper.SetDefault("host", "127.0.0.1")
	viper.SetDefault("port", 3306)
	viper.SetDefault("username", "root")
	viper.SetDefault("password", "")
	viper.SetDefault("database", "dbs")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
