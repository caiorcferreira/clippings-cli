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


