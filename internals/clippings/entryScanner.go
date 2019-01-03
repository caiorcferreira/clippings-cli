package clippings

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type EntryScanner interface {
	Scan(filePath string) ([]string, error)
}

type DefaultEntryScanner struct {}

func (s DefaultEntryScanner) Scan(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Failed to open file at '%s'\n%#v", filePath, err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(splitEntries)

	var rawEntries []string

	for scanner.Scan() {
		rawEntries = append(rawEntries, scanner.Text())
	}

	return rawEntries, nil
}

const entrySeparator = "=========="

func splitEntries(data []byte, atEOF bool) (advance int, token []byte, err error) {

	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if atEOF {
		return len(data), data, nil
	}

	if i := strings.Index(string(data), entrySeparator); i >= 0 {
		nextPos := i + len(entrySeparator)
		return nextPos, data[0:nextPos], nil
	}

	return 0, nil, nil
}
