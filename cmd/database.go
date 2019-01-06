package cmd

import (
	"github.com/caiorcferreira/kindle-clipping-cli/internals/clippings"
	"github.com/spf13/cobra"
)

var createDatabaseOutputFile string
var queryDatabaseOutputFile string

func init() {
	databaseCreateCmd.Flags().StringVarP(&createDatabaseOutputFile, "output-file", "o", "", "file to write the database")
	databaseQueryCmd.Flags().StringVarP(&queryDatabaseOutputFile, "output-file", "o", "", "file to write the database")

	databaseCmd.AddCommand(databaseCreateCmd)
	databaseCmd.AddCommand(databaseQueryCmd)
	rootCmd.AddCommand(databaseCmd)
}

var databaseCmd = &cobra.Command{
	Use: "database",
	Short: "",
}

var databaseCreateCmd = &cobra.Command{
	Use: "create",
	Run: func(cmd *cobra.Command, args []string) {
		app := clippings.NewApp()

		app.CreateDatabaseCommandRunner([]interface{}{createDatabaseOutputFile}, args)
	},
}

var databaseQueryCmd = &cobra.Command{
	Use: "query",
	Run: func(cmd *cobra.Command, args []string) {
		app := clippings.NewApp()

		app.QueryDatabaseCommandRunner([]interface{}{queryDatabaseOutputFile}, args)
	},
}
