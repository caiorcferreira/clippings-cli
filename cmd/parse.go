package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/caiorcferreira/kindle-clipping-cli-v2/internals/clippings"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(parseCmd)
}

var parseCmd = &cobra.Command{
	Use: "parse",
	Short: "Parse clippings to JSON schema",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]

		scanner := clippings.DefaultEntryScanner{}

		entries, err := clippings.Parse(scanner, filePath)

		if err != nil {
			fmt.Printf("Command failed: %#v", err)
		}

		bytes, jsonErr := json.MarshalIndent(map[string][]clippings.Entry{
			"entries": entries,
		}, "", "\t")

		if jsonErr != nil {
			fmt.Printf("Command failed: %#v", jsonErr)
		}

		os.Stdout.Write(bytes)
	},
}
