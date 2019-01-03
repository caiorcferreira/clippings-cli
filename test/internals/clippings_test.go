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
	t.Run("parse clippings from one document", func(t *testing.T) {
		filePath := "/test/file"

		entryScanner := &StubScanner{
			ExpectedRawEntries:[]string{
				`Becoming Functional (Joshua Backfield)
- Seu destaque ou posição 2046-2046 | Adicionado: quarta-feira, 14 de fevereiro de 2018 11:57:33

Local variables do not change.
==========`,
				`Becoming Functional (Joshua Backfield)
- Seu destaque ou posição 2046-2046 | Adicionado: quarta-feira, 14 de fevereiro de 2018 11:57:37

Global variables can change only references.
==========`,
			},
			ExpectedErr: nil,
		}

		entries, _ := clippings.Parse(entryScanner, filePath)

		expectedFirstEntry := clippings.Entry{
			Document: "Becoming Functional",
			Author: "Joshua Backfield",
			Kind: "highlight",
			Position: "2046-2046",
			Date: "quarta-feira, 14 de fevereiro de 2018 11:57:33",
			Content: "Local variables do not change.",
		}

		assert.Equals(t, expectedFirstEntry, entries[0])
	})

	t.Run("parse clippings fail due to non-existent file path", func(t *testing.T) {
		entryScanner := &StubScanner{
			ExpectedRawEntries:nil,
			ExpectedErr:nil,
		}

		_, err := clippings.Parse(entryScanner, "")

		assert.Equals(t, clippings.ErrFilePathEmpty, err)
	})

	t.Run("parse clippings fail due to scan error", func(t *testing.T) {
		entryScanner := &StubScanner{
			ExpectedRawEntries:nil,
			ExpectedErr: clippings.ErrScanFailed,
		}

		_, err := clippings.Parse(entryScanner, "/test/file")

		assert.Equals(t, entryScanner.ExpectedErr, err)
	})
}


