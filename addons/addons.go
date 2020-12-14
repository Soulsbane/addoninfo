package addon

// Collection A collection of Addon objects
type Collection struct {
	addons []Addon
}

// Count Get the number of intstalled addons
func (collection Collection) Count() int {
	return len(collection.addons)
}

// NewAddons create a new addon instance
func NewAddons() Collection {
	var addons Collection

	return addons
}
