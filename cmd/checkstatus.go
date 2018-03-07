package cmd

import (
	"github.com/c88lopez/dbs/src/errors"
	"github.com/c88lopez/dbs/src/statesdiff"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// checkStatusCmd represents the checkStatus command
var checkStatusCmd = &cobra.Command{
	Use:   "checkStatus",
	Short: "List statuses, marking the current database status (if any).",
	Long:  "List statuses, marking the current database status (if any).",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		initConfig()

		err := viper.ReadInConfig()
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err = errors.ErrNotInitialized
		}

		return err
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return statesdiff.CheckStatus()
	},
}

func init() {
	rootCmd.AddCommand(checkStatusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkStatusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkStatusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
