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
	Page string `json:"page"`
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

func Parse(rawClippings []string) []Entry {
	var entries []Entry

	for _, rawEntry := range rawClippings {
		entries = append(entries, parseEntry(rawEntry))
	}

	return entries
}

var (
	documentRegex = regexp.MustCompile(`^([^\-].+?)(?:\(|\n)`)
	authorRegex = regexp.MustCompile(`\(([^)]+)\)`)
	kindRegex = regexp.MustCompile(`(?:Seu|Sua) (.+?) (?:ou|na)`)
	positionRegex = regexp.MustCompile(`posição(.+?)\|`)
	pageRegex = regexp.MustCompile(`página(.+?)\|`)
	dateRegex = regexp.MustCompile(`Adicionado:(.+?)\n`)
	contentRegex = regexp.MustCompile(`(.+?)\n==========$`)
)

func parseEntry(rawClipping string) Entry {
	document := findField(documentRegex, rawClipping)
	author := findField(authorRegex, rawClipping)
	kind := findField(kindRegex, rawClipping)
	position := findField(positionRegex, rawClipping)
	page := findField(pageRegex, rawClipping)
	date := findField(dateRegex, rawClipping)
	content := findField(contentRegex, rawClipping)

	return Entry{
		Document: document,
		Author: author,
		Kind: NewKind(kind),
		Position: position,
		Page: page,
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


