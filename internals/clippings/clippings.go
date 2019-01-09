package clippings

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
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

func (a App) UpdateDatabseCommandRunner(flags []interface{}, args []string) {
	databasePath := args[0]
	clippingsFilePath := args[1]

	databaseJson, databaseErr := ioutil.ReadFile(databasePath)
	checkError(databaseErr)

	queryResult := gjson.Get(string(databaseJson), "entries")

	entries := getEntriesFromQuery(queryResult)
	entriesIds := getEntriesIds(entries)
	existingEntriesSet := makeSet(entriesIds)

	rawClippings, scannerErr := a.scanner.Scan(clippingsFilePath)
	checkError(scannerErr)

	availableEntries := Parse(rawClippings)

	for _, availableEntry := range availableEntries {
		if !existingEntriesSet[availableEntry.Id] {
			entries = append(entries, availableEntry)
		}
	}

	response := map[string][]Entry{"entries": entries,}

	jsonBytes, jsonErr := json.MarshalIndent(response, "", "\t")
	checkError(jsonErr)

	w, writeErr := makeWriter(databasePath)
	defer w.Close()

	checkError(writeErr)

	w.Write(jsonBytes)
}

func getEntriesIds(entries []Entry) (ids []string) {
	for _, entry := range entries {
		ids = append(ids, entry.Id)
	}

	return ids
}

func getEntriesFromQuery(result gjson.Result) []Entry {
	var entries []Entry

	result.ForEach(func(key, value gjson.Result) bool {

		entryMap := value.Value().(map[string]interface{})

		entries = append(entries, Entry{
			Id:       entryMap["id"].(string),
			Document: entryMap["document"].(string),
			Author:   entryMap["author"].(string),
			Kind:     NewKind(entryMap["kind"].(string)),
			Date:     entryMap["date"].(string),
			Content:  entryMap["content"].(string),
			Page:     fmt.Sprintf("%v", entryMap["page"]),
			Position: fmt.Sprintf("%v", entryMap["position"]),
		})
		return true
	})

	return entries
}

func makeSet(input []string) map[string]bool {
	set := make(map[string]bool)
	for _, inputValue := range input {
		set[inputValue] = true
	}

	return set
}

func (a App) QueryDatabaseCommandRunner(flags []interface{}, args []string) {
	query := args[0]
	databasePath := args[1]
	outputFile := flags[0].(string)

	databaseJson, databaseErr := ioutil.ReadFile(databasePath)
	checkError(databaseErr)

	queryResult := gjson.Get(string(databaseJson), query)

	var entries []interface{}
	queryResult.ForEach(func(key, value gjson.Result) bool {
		entries = append(entries, value.Value())
		return true
	})

	//TODO: change response format to use "result"
	response := map[string][]interface{}{"entries": entries,}

	jsonBytes, jsonErr := json.MarshalIndent(response, "", "\t")
	checkError(jsonErr)

	writer, writeErr := makeWriter(outputFile)
	checkError(writeErr)

	writer.Write(jsonBytes)
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

