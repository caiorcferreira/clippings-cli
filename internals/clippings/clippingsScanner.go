package clippings

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"
)

var ErrFilePathEmpty = errors.New("Argument Error: filePath must be provided")
var ErrScanFailed = errors.New("Runtime Error: entries scan failed")

type EntryScanner interface {
	Scan(filePath string) ([]string, error)
}

type DefaultEntryScanner struct {}

func (s DefaultEntryScanner) Scan(filePath string) ([]string, error) {
	if filePath == "" {
		return nil, ErrFilePathEmpty
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Failed to open file at '%s'\n%#v", filePath, err)
		return nil, ErrScanFailed
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
