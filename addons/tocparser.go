package addons

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// TocParser Used for parsing WoW TOC files.
type TocParser struct {
	values map[string]string
	files  []string
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

			// Creats a pair from example "Author: Soulsbane"
			if len(values) == 2 {
				key := strings.Trim(string(values[0]), " ")
				value := strings.Trim((values[1]), " ")

				parser.values[key] = value
			}
			// Line is a comment
		} else if len(line) == 0 || (strings.HasPrefix(line, "##") && !strings.Contains(line, ":")) {
			continue
			// Line is a empty or a filename. If blank ignore.
		} else {
			if strings.TrimSpace(line) != "" {
				parser.files = append(parser.files, line)
			}
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

// GetAuthor Get the addons author name.
func (parser TocParser) GetAuthor() string {
	return parser.GetEntry("Author")
}

// GetVersion Get the addons version.
func (parser TocParser) GetVersion() string {
	return parser.GetEntry("Version")
}

// GetTitle Get the addons title.
func (parser TocParser) GetTitle() string {
	return parser.GetEntry("Title")
}

// GetNotes Get the addons notes.
func (parser TocParser) GetNotes() string {
	return parser.GetEntry("Notes")
}

// GetInterface Get the addons interface version.
func (parser TocParser) GetInterface() string {
	return parser.GetEntry("Interface")
}

// Dump dumps the key/value pairs to stdout
func (parser TocParser) Dump() {
	for key, value := range parser.values {
		fmt.Printf("%s => %s\n", key, value)
	}
}