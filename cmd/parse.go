package cmd

import (
	"github.com/caiorcferreira/kindle-clipping-cli/internals/clippings"
	"github.com/spf13/cobra"
)

var parseOutputFile string

func init() {
	parseCmd.Flags().StringVarP(&parseOutputFile, "output-file", "o", "", "file to write the parsing result")
	rootCmd.AddCommand(parseCmd)
}

var parseCmd = &cobra.Command{
	Use: "parse",
	Short: "Parse clippings to JSON schema",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app := clippings.NewApp()

		app.ParseCommandRunner([]interface{}{parseOutputFile}, args)
	},
}
