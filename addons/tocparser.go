package addon

import (
	"fmt"
)

// TocParser Used for parsing WoW TOC files.
type TocParser struct {
	values map[string]string
}

// NewTocParser Creates a new TocParser
func NewTocParser() TocParser {
	var parser TocParser

	parser.values = make(map[string]string)
	return parser
}

// AddEntry Adds a new key/value pair
func (parser TocParser) AddEntry(key string, value string) {
	parser.values[key] = value
}

// HasEntry Check if an entry exists.
func (parser TocParser) HasEntry(name string) bool {
	if _, found := parser.values[name]; found {
		return true
	}

	return false
}

// GetEntry Get an entry
func (parser TocParser) GetEntry(name string) string {
	if value, found := parser.values[name]; found {
		return value
	}

	return ""
}

// Dump dumps the key/value pairs to stdout
func (parser TocParser) Dump() {
	for key, value := range parser.values {
		fmt.Printf("%s => %s\n", key, value)
	}
}
