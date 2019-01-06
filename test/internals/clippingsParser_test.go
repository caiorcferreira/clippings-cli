package clippings

import (
	"github.com/caiorcferreira/kindle-clipping-cli-v2/internals/clippings"
	assert "github.com/caiorcferreira/kindle-clipping-cli-v2/test/util"
	"testing"
)

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
				Id: "6c49cc5682404f64fde91f9e06e4ad54",
				Document: "Becoming Functional",
				Author: "Joshua Backfield",
				Kind: clippings.Highlight,
				Position: "2046-2046",
				Date: "quarta-feira, 14 de fevereiro de 2018 11:57:33",
				Content: "Local variables do not change.",
			},
			{
				Id: "68efb25a0a45ace422a2d5b7dd60b150",
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
			Id: "2e865ac9f6fca2f245c366918c8ab6d3",
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
			Id: "99dd7706b2644352778c613df155a717",
			Document: "Building Microservices",
			Author:   "Sam Newman",
			Kind:     clippings.Highlight,
			Page: "227-227",
			Date:     "terça-feira, 4 de dezembro de 2018 19:45:12",
			Content:  "REST In Practice",
		}

		assert.Equals(t, entry, entries[0])
	})

	t.Run("parse clipping without author", func(t *testing.T) {
		rawClippings := []string{
			`Building Microservices - Sam Newman  
- Seu destaque na página 212-212 | Adicionado: sábado, 1 de dezembro de 2018 23:48:59

Michael Nygard’s book Release It
==========`,
		}

		entries := clippings.Parse(rawClippings)

		entry := clippings.Entry{
			Id: "b9ec65e8c60375b9c5a511940aaafd4c",
			Document: "Building Microservices - Sam Newman",
			Author: "",
			Kind: clippings.Highlight,
			Page: "212-212",
			Date: "sábado, 1 de dezembro de 2018 23:48:59",
			Content: "Michael Nygard’s book Release It",
		}

		assert.Equals(t, entry, entries[0])
	})

	t.Run("parse clipping without author", func(t *testing.T) {
		rawClippings := []string{
			`A Cura de Schopenhauer (Irvin D. Yalom)
- Seu marcador na página 13 | posição 188 | Adicionado: domingo, 7 de junho de 2015 18:21:16


==========`,
		}

		entries := clippings.Parse(rawClippings)

		entry := clippings.Entry{
			Id: "e0b3f36727107f0b57777404632709c6",
			Document: "A Cura de Schopenhauer",
			Author: "Irvin D. Yalom",
			Kind: clippings.Bookmark,
			Page: "13",
			Position: "188",
			Date: "domingo, 7 de junho de 2015 18:21:16",
			Content: "",
		}

		assert.Equals(t, entry, entries[0])
	})
}


