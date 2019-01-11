package clippings

import (
	"encoding/json"
	"github.com/caiorcferreira/kindle-clipping-cli/internals/clippings"
	assert "github.com/caiorcferreira/kindle-clipping-cli/test/util"
	"testing"
)

type StubFileReader struct {
	expectedResult []byte
}

func (s StubFileReader) ReadFile() ([]byte, error) {
	return s.expectedResult, nil
}


func TestEntriesRepository(t *testing.T) {
	t.Run("get all entries from JSON database", func(t *testing.T) {
		stubEntries := []clippings.Entry{
			{
				Id: "6c49cc5682404f64fde91f9e06e4ad54",
				Document: "Becoming Functional",
				Author: "Joshua Backfield",
				Kind: clippings.Highlight,
				Position: "2046-2046",
				Date: "quarta-feira, 14 de fevereiro de 2018 11:57:33",
				Content: "Local variables do not change.",
			},
		}

		marshalEntry, _ := json.Marshal(map[string][]clippings.Entry{"entries":stubEntries})
		reader := StubFileReader{marshalEntry}
		repository := &clippings.EntriesRepository{Reader: reader}

		entries := repository.GetAll()

		assert.Equals(t, stubEntries, entries)
	})
	
	t.Run("query entries from JSON database", func(t *testing.T) {
		stubEntries := []map[string]interface{}{
			{
				"id": "6c49cc5682404f64fde91f9e06e4ad54",
				"document": "Becoming Functional",
				"author": "Joshua Backfield",
				"kind": clippings.Highlight.String(),
				"position": "2046-2046",
				"date": "quarta-feira, 14 de fevereiro de 2018 11:57:33",
				"content": "Local variables do not change.",
			},
			{
				"id": "6c49cc5682404f64fde91f9e06e4ad54",
				"document": "Becoming Functional",
				"author": "Joshua Backfield",
				"kind": clippings.Note.String(),
				"position": "2046-2046",
				"date": "quarta-feira, 14 de fevereiro de 2018 11:57:33",
				"content": "Local variables do not change.",
			},
		}

		file := map[string]interface{}{"entries": stubEntries}
		marshalEntries, _ := json.Marshal(file)
		reader := StubFileReader{marshalEntries}
		repository := &clippings.EntriesRepository{Reader: reader}

		query := "entries.#[kind==\"note\"]#"
		entries := repository.GetByQuery(query)

		assert.Equals(t, stubEntries[1], entries[0])
	})
}
