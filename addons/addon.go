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

// GetVersion Get the addons version.
func (addon Addon) GetVersion() string {
	return addon.parser.GetEntry("Version")
}

// GetTitle Get the addons title.
func (addon Addon) GetTitle() string {
	return addon.parser.GetEntry("Title")
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
	addon.parser.Dump()
	fmt.Println("HasEntry for Version: ", addon.parser.HasEntry("Version"))
	fmt.Println("HasEntry for VersionSSSSSSSSS: ", addon.parser.HasEntry("VersionSSSSSSSSS"))
	fmt.Println(addon.parser.GetEntry("Author"))
	fmt.Println(addon.parser.GetEntry("AuthorZZZZ"), "<<<")
}
