package addons

import (
	"fmt"
	"io/ioutil"
	"strings"
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

// ParseString Parses content into an associative array
func (parser *TocParser) ParseString(content string) {
	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		line := strings.TrimSpace(line)

		if strings.HasPrefix(line, "##") {
			line := strings.Trim(line, "#")
			values := strings.Split(line, ":")

			// Creats a pair from example "Key: Value"
			if len(values) == 2 {
				key := strings.Trim(string(values[0]), " ")
				value := strings.Trim((values[1]), " ")

				parser.values[key] = value
			}
			// Line is a comment
		} else if len(line) == 0 || (strings.HasPrefix(line, "##") && !strings.Contains(line, ":")) {
			continue
			// Line is a file. If blank ignore.
		} else {

		}
	}

}

// ParseFile Loads a TOC file's contents into a string and calls ParseString
func (parser TocParser) ParseFile(fileName string) {
	content, err := ioutil.ReadFile(fileName)

	if err != nil {
		parser.ParseString(string(content))
	}
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
