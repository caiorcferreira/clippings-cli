package clippings

import (
	"github.com/caiorcferreira/kindle-clipping-cli-v2/internals/clippings"
	assert "github.com/caiorcferreira/kindle-clipping-cli-v2/test/util"
	"testing"
)

func TestScanClippings(t *testing.T) {
	t.Run("scan clippings from file", func(t *testing.T) {
		filePath := "/Users/caioferreira/sandbox/kindle-clippings-cli-v2/test/data/clippings.txt"

		scanner := clippings.DefaultEntryScanner{}

		rawRntries, _ := scanner.Scan(filePath)

		expectedFirstEntry := 	`Becoming Functional (Joshua Backfield)
- Seu destaque ou posição 2046-2046 | Adicionado: quarta-feira, 14 de fevereiro de 2018 11:57:33

Local variables do not change.
==========`

		assert.Equals(t, expectedFirstEntry, rawRntries[0])
	})
}
