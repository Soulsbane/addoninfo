package addons

import (
	"fmt"
	"strings"
)

// Collection A collection of Addon objects
type Collection struct {
	addons []Addon
}

// NewCollection create a new addon instance
func NewCollection() Collection {
	var addons Collection

	return addons
}

// Count Get the number of intstalled addons
func (collection Collection) Count() int {
	return len(collection.addons)
}

// Add Add an Addon to the collection
func (collection *Collection) Add(name string, addon Addon) {
	collection.addons = append(collection.addons, addon)
}

// GetAddon Get an instance of an addon.
func (collection Collection) GetAddon(name string) (Addon, error) {
	for _, addon := range collection.addons {
		if strings.ToLower(addon.GetTitle()) == strings.ToLower(name) {
			return addon, nil
		}
	}

	return Addon{}, fmt.Errorf("addon name does not exist")
}
