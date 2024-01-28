package addons

import (
	"fmt"
	"github.com/Soulsbane/tocparser/tocparser"
	"os"
	"regexp"
	"strings"
)

// Addon the addon instance
type Addon struct {
	parser  tocparser.Parser
	dirName string // If Title TOC field is missing or blank just use the addons directory name
}

// NewAddon create a new addon instance
func NewAddon(dirName string, tocFileName string) Addon {
	var addon Addon

	addon.parser = tocparser.New()
	addon.dirName = dirName
	err := addon.parser.LoadFile(tocFileName)

	if err != nil {
		fmt.Println("Error loading TOC file: ", err)
	}

	return addon
}

// GetAuthor Get the name of the author
func (addon Addon) GetAuthor() string {
	return addon.parser.GetEntry("Author")
}

// GetVersion Get the addons version.
func (addon Addon) GetVersion() string {
	return addon.parser.GetEntry("Version")
}

// GetTitle Get the addons title.
func (addon Addon) GetTitle() string {
	name := addon.parser.GetEntry("Title")

	if len(name) == 0 || name == " " {
		return addon.dirName
	}

	/*
		Some addon authors like to colorize the title text. Remove it here
		Kui |cff9966ffNameplates
		Kui |cff9966ffNameplates:|r |cffffffffCore|r
	*/
	if strings.Contains(name, "|") {
		re := regexp.MustCompile(`\|c\w{8}`) // Remove |cff9966ff like strings
		subbed := string(re.ReplaceAll([]byte(name), []byte("")))

		return strings.ReplaceAll(subbed, "|r", "") // Finally remove any |r from the title
	}

	return name
}

// GetDirName The the name of the directory the addon is stored on.
func (addon Addon) GetDirName() string {
	return addon.dirName
}

// GetNotes Get the addons notes.
func (addon Addon) GetNotes() string {
	return addon.parser.GetEntry("Notes")
}

// GetInterface Get the addons interface version.
func (addon Addon) GetInterface() string {
	return addon.parser.GetEntry("Interface")
}

// TestParser Test the parser.
func (addon Addon) TestParser() {
	addon.parser.AddEntry("Version", "1.0")
	addon.parser.AddEntry("Author", "Soulsbane")
	addon.parser.AddEntry("Name", "TocParser")
	addon.parser.DumpEntries(os.Stdout)
	fmt.Println("HasEntry for Version: ", addon.parser.HasEntry("Version"))
	fmt.Println("HasEntry for VersionSSSSSSSSS: ", addon.parser.HasEntry("VersionSSSSSSSSS"))
	fmt.Println(addon.parser.GetEntry("Author"))
	fmt.Println(addon.parser.GetEntry("AuthorZZZZ"), "<<<")
}
