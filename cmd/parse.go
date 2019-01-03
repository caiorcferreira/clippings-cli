package cmd

import (
	"github.com/caiorcferreira/kindle-clipping-cli-v2/internals/clippings"
	"github.com/spf13/cobra"
)

var outputFile string

func init() {
	parseCmd.Flags().StringVarP(&outputFile, "output-file", "o", "", "file to write the parsing result")
	rootCmd.AddCommand(parseCmd)
}

var parseCmd = &cobra.Command{
	Use: "parse",
	Short: "Parse clippings to JSON schema",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app := clippings.App{}
		app.ParseCommandRunner([]interface{}{outputFile}, args)
	},
}
