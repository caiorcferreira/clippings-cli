package clippings

import (
	"encoding/json"
	"log"
	"os"
)

type App struct {
	scanner Scanner
}

func NewApp() App {
	scanner := DefaultScanner{}
	app := App{scanner}

	return app
}

func (a App) ParseCommandRunner(flags []interface{}, args []string) {
	clippingsFilePath := args[0]
	outputFile := flags[0].(string)

	rawClippings, scannerErr := a.scanner.Scan(clippingsFilePath)
	checkError(scannerErr)

	entries := Parse(rawClippings)
	response := map[string][]Entry{"entries": entries,}

	jsonBytes, jsonErr := json.MarshalIndent(response, "", "\t")
	checkError(jsonErr)

	w, writeErr := makeWriter(outputFile)
	defer w.Close()

	checkError(writeErr)

	w.Write(jsonBytes)
}

func (a App) CreateDatabaseCommandRunner(flags []interface{}, args []string) {
	clippingsFilePath := args[0]
	outputFile := flags[0].(string)
	if outputFile == "" {
		outputFile = "./database.json"
	}

	rawClippings, scannerErr := a.scanner.Scan(clippingsFilePath)
	checkError(scannerErr)

	entries := Parse(rawClippings)
	response := map[string][]Entry{"entries": entries,}

	jsonBytes, jsonErr := json.MarshalIndent(response, "", "\t")
	checkError(jsonErr)

	w, writeErr := makeWriter(outputFile)
	defer w.Close()

	checkError(writeErr)

	w.Write(jsonBytes)
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Command failed: %#v", err)
	}
}

func makeWriter(outputFile string) (*os.File, error) {
	if outputFile == "" {
		return os.Stdout, nil
	} else {
		return os.Create(outputFile)
	}
}

