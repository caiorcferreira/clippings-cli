package clippings

import (
	"github.com/caiorcferreira/kindle-clipping-cli-v2/internals/clippings"
	assert "github.com/caiorcferreira/kindle-clipping-cli-v2/test/util"
	"testing"
)

type StubScanner struct {
	ExpectedRawEntries []string
	ExpectedErr        error
}

func (s *StubScanner) Scan(filePath string) ([]string, error) {
	return s.ExpectedRawEntries, s.ExpectedErr
}

func TestParseClippingsFile(t *testing.T) {
	t.Run("parse clippings", func(t *testing.T) {
		rawEntries := []string{
				`Becoming Functional (Joshua Backfield)
- Seu destaque ou posição 2046-2046 | Adicionado: quarta-feira, 14 de fevereiro de 2018 11:57:33

Local variables do not change.
==========`,
				`Becoming Functional (Joshua Backfield)
- Seu destaque ou posição 2046-2046 | Adicionado: quarta-feira, 14 de fevereiro de 2018 11:57:37

Global variables can change only references.
==========`,
		}

		entries := clippings.Parse(rawEntries)

		expectedEntries := []clippings.Entry{
			{
				Document: "Becoming Functional",
				Author: "Joshua Backfield",
				Kind: clippings.Highlight,
				Position: "2046-2046",
				Date: "quarta-feira, 14 de fevereiro de 2018 11:57:33",
				Content: "Local variables do not change.",
			},
			{
				Document: "Becoming Functional",
				Author: "Joshua Backfield",
				Kind: clippings.Highlight,
				Position: "2046-2046",
				Date: "quarta-feira, 14 de fevereiro de 2018 11:57:37",
				Content: "Global variables can change only references.",
			},
		}

		assert.Equals(t, expectedEntries, entries)
	})

	t.Run("parse clipping with note", func(t *testing.T) {
		rawClippings := []string{
			`A Cura de Schopenhauer (Irvin D. Yalom)
- Sua nota ou posição 939 | Adicionado: segunda-feira, 8 de junho de 2015 17:34:04

Esta citação é de Hobbes
==========`,
		}

		entries := clippings.Parse(rawClippings)

		entry := clippings.Entry{
			Document: "A Cura de Schopenhauer",
			Author:   "Irvin D. Yalom",
			Kind:     clippings.Note,
			Position: "939",
			Date:     "segunda-feira, 8 de junho de 2015 17:34:04",
			Content:  "Esta citação é de Hobbes",
		}

		assert.Equals(t, entry, entries[0])
	})

	t.Run("parse clipping with page count instead of position", func(t *testing.T) {
		rawClippings := []string{
			`Building Microservices (Sam Newman)
- Seu destaque na página 227-227 | Adicionado: terça-feira, 4 de dezembro de 2018 19:45:12

REST In Practice
==========`,
		}

		entries := clippings.Parse(rawClippings)

		entry := clippings.Entry{
			Document: "Building Microservices",
			Author:   "Sam Newman",
			Kind:     clippings.Highlight,
			Page: "227-227",
			Date:     "terça-feira, 4 de dezembro de 2018 19:45:12",
			Content:  "REST In Practice",
		}

		assert.Equals(t, entry, entries[0])
	})

	//t.Run("parse clippings fail due to non-existent file path", func(t *testing.T) {
	//	entryScanner := &StubScanner{
	//		ExpectedRawEntries:nil,
	//		ExpectedErr:nil,
	//	}
	//
	//	_, err := clippings.Parse(entryScanner, "")
	//
	//	assert.Equals(t, clippings.ErrFilePathEmpty, err)
	//})

	//t.Run("parse clippings fail due to scan error", func(t *testing.T) {
	//	entryScanner := &StubScanner{
	//		ExpectedRawEntries:nil,
	//		ExpectedErr: clippings.ErrScanFailed,
	//	}
	//
	//	_, err := clippings.Parse(entryScanner, "/test/file")
	//
	//	assert.Equals(t, entryScanner.ExpectedErr, err)
	//})
}


