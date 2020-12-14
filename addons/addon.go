package addons

import "fmt"

// Addon the addon instance
type Addon struct {
	parser TocParser
}

// NewAddon create a new addon instance
func NewAddon() Addon {
	var addon Addon
	addon.parser = NewTocParser()

	return addon
}

// GetAuthor Get the name of the author
func (addon Addon) GetAuthor() string {
	return addon.parser.GetEntry("Author")
}

// TestParser Test the parser.
func (addon Addon) TestParser() {
	addon.parser.AddEntry("Version", "1.0")
	addon.parser.AddEntry("Author", "Soulsbane")
	addon.parser.AddEntry("Name", "TocParser")
	addon.parser.Dump()
	fmt.Println("HasEntry for Version: ", addon.parser.HasEntry("Version"))
	fmt.Println("HasEntry for VersionSSSSSSSSS: ", addon.parser.HasEntry("VersionSSSSSSSSS"))
	fmt.Println(addon.parser.GetEntry("Author"))
	fmt.Println(addon.parser.GetEntry("AuthorZZZZ"), "<<<")
}
