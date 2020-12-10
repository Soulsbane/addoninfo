package main

import "fmt"

// TocParser Used for parsing WoW TOC files.
type TocParser struct {
	values map[string]string
}

// NewTocParser Creates a new TocParser
func NewTocParser() *TocParser {
	var parser TocParser

	parser.values = make(map[string]string)
	return &parser
}

// AddEntry Adds a new key/value pair
func (parser *TocParser) AddEntry(key string, value string) {
	parser.values[key] = value
}

// Dump dumps the key/value pairs to stdout
func (parser *TocParser) Dump() {
	for key, value := range parser.values {
		fmt.Printf("%s => %s\n", key, value)
	}
}
