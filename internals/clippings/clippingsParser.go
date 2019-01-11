package clippings

import (
	"crypto/md5"
	"encoding/hex"
	"regexp"
	"strings"
)

type Entry struct {
	Id string `json:"id"`
	Document string	`json:"document"`
	Author string `json:"author"`
	Kind Kind `json:"kind"`
	Position string `json:"position,omitempty"`
	Page string `json:"page,omitempty"`
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
	kindsFromClippings := map[string]Kind{
		"destaque": Highlight,
		"nota": Note,
		"marcador": Bookmark,
	}

	rawKinds := map[string]Kind{
		"highlight": Highlight,
		"note": Note,
		"bookmark": Bookmark,
	}

	kind, ok := kindsFromClippings[value]

	if ok {
		return kind
	} else {
		return rawKinds[value]
	}
}

func Parse(rawClippings []string) []Entry {
	var entries []Entry

	for _, rawEntry := range rawClippings {
		if strings.TrimSpace(rawEntry) != "" {
			entries = append(entries, parseEntry(rawEntry))
		}
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

	id := generateId(rawClipping)

	return Entry{
		Id: id,
		Document: document,
		Author: author,
		Kind: NewKind(kind),
		Position: position,
		Page: page,
		Date: date,
		Content: content,
	}
}

func generateId(target string) string {
	encoder := md5.New()
	encoder.Write([]byte(target))

	return hex.EncodeToString(encoder.Sum(nil))
}

func findField(r *regexp.Regexp, target string) string {
	submatch := r.FindStringSubmatch(strings.TrimSpace(target))

	if len(submatch) < 2 {
		return ""
	}

	return strings.TrimSpace(submatch[1])
}


