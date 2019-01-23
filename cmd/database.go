package cmd

import (
	"github.com/caiorcferreira/kindle-clipping-cli/internals/clippings"
	"github.com/spf13/cobra"
)

var createDatabaseOutputFile string
var queryDatabaseOutputFile string

func init() {
	databaseCreateCmd.Flags().StringVarP(&createDatabaseOutputFile, "output-file", "o", "", "file to write the created database")
	databaseQueryCmd.Flags().StringVarP(&queryDatabaseOutputFile, "output-file", "o", "", "file to write the query results")

	databaseCmd.AddCommand(databaseCreateCmd)
	databaseCmd.AddCommand(databaseUpdateCmd)
	databaseCmd.AddCommand(databaseQueryCmd)
	rootCmd.AddCommand(databaseCmd)
}

var databaseCmd = &cobra.Command{
	Use: "database",
	Short: "Manage the JSON Clippings database",
}

var databaseCreateCmd = &cobra.Command{
	Use:                    "create",
	Short:                  "Create a new JSON database from a raw clippings file",
	Example:                "clippings database create [-o DATABASE_PATH] <CLIPPINGS_PATH>",
	Run: func(cmd *cobra.Command, args []string) {
		app := clippings.NewApp()

		app.CreateDatabaseCommandRunner([]interface{}{createDatabaseOutputFile}, args)
	},
}

var databaseUpdateCmd = &cobra.Command{
	Use: "update",
	Short: "Updates an existing database with raw clippings from a file",
	Long: `Updates a JSON database with raw clippings from a file.
			The tool uses a simple hashing algorithm in order to avoid inserting clippings that already exists in the database.
			But notice that sometimes Kindle introduces duplicated clippings that cannot be avoided be this mechanism.`,
	Example: "clippings database update <DATABASE_PATH> <CLIPPINGS_PATH>",
	Run: func(cmd *cobra.Command, args []string) {
		app := clippings.NewApp()

		app.UpdateDatabseCommandRunner([]interface{}{}, args)
	},
}

var databaseQueryCmd = &cobra.Command{
	Use: "query",
	Short: "Executes a query against a JSON database",
	Long: `Executes a query using path syntax for search against a JSON database, optionally writing the query results to a given file.
			To see more about the path syntax refer to https://github.com/tidwall/gjson#path-syntax`,
	Example: "clippings database query [-o OUTPUT_FILE] <QUERY> <DATABASE_PATH>",
	Run: func(cmd *cobra.Command, args []string) {
		app := clippings.NewApp()

		app.QueryDatabaseCommandRunner([]interface{}{queryDatabaseOutputFile}, args)
	},
}
