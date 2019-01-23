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
	reader *DefaultFileReader
	repository *EntriesRepository
}

func NewApp() App {
	scanner := DefaultScanner{}
	reader := &DefaultFileReader{}
	repository := &EntriesRepository{Reader:reader}

	app := App{scanner, reader, repository}

	return app
}

type DefaultFileReader struct {
	FilePath string
}

func (f *DefaultFileReader) ReadFile() ([]byte, error) {
	return ioutil.ReadFile(f.FilePath)
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

	a.reader.FilePath = databasePath
	existingEntries := a.repository.GetAll()
	existingIdsSet := makeSet(getEntriesIds(existingEntries))

	rawClippings, scannerErr := a.scanner.Scan(clippingsFilePath)
	checkError(scannerErr)

	availableEntries := Parse(rawClippings)

	for _, availableEntry := range availableEntries {
		if !existingIdsSet[availableEntry.Id] {
			existingEntries = append(existingEntries, availableEntry)
		}
	}

	response := map[string][]Entry{"existingEntries": existingEntries,}

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
			Date:     entryMap["date"].(string),
			Content:  entryMap["content"].(string),
			Kind:     NewKind(entryMap["kind"].(string)),
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

	a.reader.FilePath = databasePath
	entries := a.repository.GetByQuery(query)

	response := map[string][]interface{}{"result": entries,}

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

