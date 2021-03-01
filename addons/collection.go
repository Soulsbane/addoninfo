package addons

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/brettski/go-termtables"
	"github.com/spf13/cast"
)

// Collection A collection of Addon objects
type Collection struct {
	addons           []Addon
	currentInterface uint
}

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

// NewCollection create a new collection instance
func NewCollection() Collection {
	var addons Collection
	addons.currentInterface = 90002

	return addons
}

func (collection *Collection) isAddonOutdated(interfaceVersion string) (bool, string) {
	if cast.ToUint(interfaceVersion) >= collection.currentInterface {
		return false, "No"
	}

	return true, "Yes"
}

// Build Creates a list of installed addons
func (collection *Collection) Build(path string) {
	dirs, err := ioutil.ReadDir(path)

	if err != nil {
		fmt.Println(err)
	}

	for _, dir := range dirs {
		if dir.IsDir() {
			tocFileName := filepath.Join(dir.Name(), dir.Name()+".toc")

			if fileExists(tocFileName) {
				addon := NewAddon(dir.Name(), tocFileName)
				collection.Add(addon)
			}
		}
	}
}

// Count Get the number of intstalled addons
func (collection Collection) Count() int {
	return len(collection.addons)
}

// Add Add an Addon to the collection
func (collection *Collection) Add(addon Addon) {
	collection.addons = append(collection.addons, addon)
}

// GetAddon Get an instance of an addon.
func (collection Collection) GetAddon(name string) (Addon, error) {
	for _, addon := range collection.addons {
		if strings.ToLower(addon.GetTitle()) == strings.ToLower(name) {
			return addon, nil
		}
	}

	return Addon{}, fmt.Errorf("%s addon could not be found", name)
}

// List Print a formatted list of installed addons.
func (collection Collection) List(command string) {
	table := termtables.CreateTable()
	table.AddHeaders("Name", "Version", "Outdated", "Interface")

	for _, addon := range collection.addons {
		title := addon.GetTitle()
		version := addon.GetVersion()
		outdated, yesNo := collection.isAddonOutdated(addon.GetInterface())

		if command == "outdated" {
			if outdated {
				table.AddRow(title, version, yesNo, addon.GetInterface())
			}
		} else {
			table.AddRow(title, version, yesNo, addon.GetInterface())
		}
	}
	if len(collection.addons) > 0 {
		fmt.Println(table.Render())
	} else {
		fmt.Println("There are no addons to list!")
	}
}
