package clippings

import (
	"errors"
	"log"
	"regexp"
	"strings"
)


var ErrFilePathEmpty = errors.New("Argument Error: filePath must be provided")
var ErrScanFailed = errors.New("Runtime Error: entries scan failed")

type Entry struct {
	Document string	`json:"document"`
	Author string `json:"author"`
	Kind string `json:"kind"`
	Position string `json:"position"`
	Date string `json:"date"`
	Content string `json:"content"`
}

func Parse(scanner EntryScanner, filePath string) ([]Entry, error) {
	if filePath == "" {
		log.Println(ErrFilePathEmpty)
		return nil, ErrFilePathEmpty
	}

	rawEntries, scanErr := scanner.Scan(filePath)

	if scanErr != nil {
		log.Println(ErrScanFailed)
		return nil, ErrScanFailed
	}

	var entries []Entry

	for _, rawEntry := range rawEntries {
		entries = append(entries, parseEntry(rawEntry))
	}

	return entries, nil
}


func parseEntry(rawEntry string) Entry {
	documentRegex := regexp.MustCompile(`(.+?)\(`)
	document := documentRegex.FindStringSubmatch(rawEntry)[1]

	authorRegex := regexp.MustCompile(`\(([^)]+)\)`)
	author := authorRegex.FindStringSubmatch(rawEntry)[1]


	kindRegex := regexp.MustCompile(`Seu (.+?) ou`)

	var kind string
	if kindRegex.FindStringSubmatch(rawEntry)[1] == "destaque" {
		kind = "highlight"
	}

	positionRegex := regexp.MustCompile(`posição(.+?)\|`)
	position := positionRegex.FindStringSubmatch(rawEntry)[1]

	dateRegex := regexp.MustCompile(`Adicionado:(.+?)\n`)
	date := dateRegex.FindStringSubmatch(rawEntry)[1]

	contentRegex := regexp.MustCompile(`(.+?)\n==========$`)
	content := contentRegex.FindStringSubmatch(rawEntry)[1]


	return Entry{
		Document:strings.TrimSpace(document),
		Author:author,
		Kind:kind,
		Position:strings.TrimSpace(position),
		Date:strings.TrimSpace(date),
		Content: content,
	}
}

