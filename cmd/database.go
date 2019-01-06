package cmd

import (
	"github.com/caiorcferreira/kindle-clipping-cli/internals/clippings"
	"github.com/spf13/cobra"
)

var databaseOutputFile string

func init() {
	databaseCreateCmd.Flags().StringVarP(&databaseOutputFile, "output-file", "o", "", "file to write the database")

	rootCmd.AddCommand(databaseCmd)
	databaseCmd.AddCommand(databaseCreateCmd)
}

var databaseCmd = &cobra.Command{
	Use: "database",
	Short: "",
}

var databaseCreateCmd = &cobra.Command{
	Use: "create",
	Run: func(cmd *cobra.Command, args []string) {
		scanner := clippings.DefaultScanner{}
		app := clippings.App{scanner}

		app.CreateDatabaseCommandRunner([]interface{}{databaseOutputFile}, args)
	},
}
