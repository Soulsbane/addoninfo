package addons

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Collection A collection of Addon objects
type Collection struct {
	addons []Addon
}

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// NewCollection create a new addon instance
func NewCollection() Collection {
	var addons Collection

	return addons
}

// Build Creates a list of installed addons found in path
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
	fmt.Println("Installed Addons:")
	for _, addon := range collection.addons {
		fmt.Println(addon.GetTitle())
	}
}
