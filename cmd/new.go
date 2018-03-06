package cmd

import (
	"github.com/c88lopez/dbs/src/database"
	"github.com/c88lopez/dbs/src/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Build and persist the current database state.",
	Long:  "Build and persist the current database state.",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		initConfig()

		err := viper.ReadInConfig()
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err = errors.ErrNotInitialized
		}

		return err
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return database.New()
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
