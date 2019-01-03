package clippings

import (
	"log"
	"regexp"
	"strings"
)



type Entry struct {
	Document string	`json:"document"`
	Author string `json:"author"`
	Kind Kind `json:"kind"`
	Position string `json:"position"`
	Date string `json:"date"`
	Content string `json:"content"`
}

type Kind string

const (
	Highlight Kind = "highlight"
	Note Kind = "note"
	Bookmark Kind = "bookmark"
)

func (k Kind) String() string {
	return string(k)
}

func NewKind(value string) Kind {
	rawKinds := map[string]Kind{
		"destaque": Highlight,
		"nota": Note,
		"marcador": Bookmark,
	}

	return rawKinds[value]
}

func Parse(rawEntries []string) []Entry {
	var entries []Entry

	for _, rawEntry := range rawEntries {
		entries = append(entries, parseEntry(rawEntry))
	}

	return entries
}

var (
	documentRegex = regexp.MustCompile(`(.+?)\(`)
	authorRegex = regexp.MustCompile(`\(([^)]+)\)`)
	kindRegex = regexp.MustCompile(`Seu (.+?) ou`)
	positionRegex = regexp.MustCompile(`posição(.+?)\|`)
	dateRegex = regexp.MustCompile(`Adicionado:(.+?)\n`)
	contentRegex = regexp.MustCompile(`(.+?)\n==========$`)
)

func parseEntry(rawEntry string) Entry {
	document := findField(documentRegex, rawEntry)
	author := findField(authorRegex, rawEntry)
	kind := findField(kindRegex, rawEntry)
	position := findField(positionRegex, rawEntry)
	date := findField(dateRegex, rawEntry)
	content := findField(contentRegex, rawEntry)

	return Entry{
		Document: document,
		Author: author,
		Kind: NewKind(kind),
		Position: position,
		Date: date,
		Content: content,
	}
}

func findField(r *regexp.Regexp, target string) string {
	submatch := r.FindStringSubmatch(target)

	if len(submatch) < 2 {
		log.Printf("Could not find any match with regex %s in the given string %s", r, target)
		return ""
	}

	return strings.TrimSpace(submatch[1])
}


