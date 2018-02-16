package cmd

import (
	"github.com/c88lopez/dbs/src/statesDiff"
	"github.com/spf13/cobra"
)

// checkStatusCmd represents the checkStatus command
var checkStatusCmd = &cobra.Command{
	Use:   "checkStatus",
	Short: "List statuses, marking the current database status (if any).",
	Long:  "List statuses, marking the current database status (if any).",
	RunE: func(cmd *cobra.Command, args []string) error {
		return statesDiff.CheckStatus()
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
