package clippings

import (
	"fmt"
	"github.com/tidwall/gjson"
	"log"
)

type FileReader interface {
	ReadFile() ([]byte, error)
}

type EntriesRepository struct {
	Reader FileReader
}

func (e *EntriesRepository) GetAll() []Entry {
	file, err := e.Reader.ReadFile()
	if err != nil {
		log.Fatal("Failed when reading file")
		return nil
	}

	result := gjson.GetBytes(file, "entries")

	return e.getEntriesFromQuery(result)
}

func (e *EntriesRepository) getEntriesFromQuery(result gjson.Result) []Entry {
	var entries []Entry

	result.ForEach(func(key, value gjson.Result) bool {

		entryMap := value.Value().(map[string]interface{})

		entries = append(entries, Entry{
			Id:       formatInput(entryMap["id"]),
			Document: formatInput(entryMap["document"]),
			Author:   formatInput(entryMap["author"]),
			Date:     formatInput(entryMap["date"]),
			Content:  formatInput(entryMap["content"]),
			Page:     formatInput(entryMap["page"]),
			Position: formatInput(entryMap["position"]),
			Kind:     NewKind(formatInput(entryMap["kind"])),
		})
		return true
	})

	return entries
}

func formatInput(input interface{}) string {
	if input == nil {
		return ""
	}

	formatted := fmt.Sprintf("%v", input)
	return formatted
}


func (e *EntriesRepository) GetByQuery(query string) []interface{} {
	file, err := e.Reader.ReadFile()
	if err != nil {
		log.Fatal("Failed when reading file")
		return nil
	}

	result := gjson.GetBytes(file, query)
	return e.parseQueryResult(result)
}

func (e *EntriesRepository) parseQueryResult(result gjson.Result) []interface{} {
	var parsed []interface{}
	result.ForEach(func(key, value gjson.Result) bool {
		parsed = append(parsed, value.Value())
		return true
	})

	return parsed
}
