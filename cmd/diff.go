package cmd

import (
	"github.com/c88lopez/dbs/src/errors"
	"github.com/c88lopez/dbs/src/statesDiff"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// diffCmd represents the diff command
var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Get queries to run",
	Long:  `Get queries to run.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		initConfig()

		err := viper.ReadInConfig()
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err = errors.ErrNotInitialized
		}

		return err
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return statesDiff.Diff()
	},
}

func init() {
	rootCmd.AddCommand(diffCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// diffCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diffCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
