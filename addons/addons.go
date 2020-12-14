package addon

// Collection A collection of Addon objects
type Collection struct {
	addons []Addon
}

// Count Get the number of intstalled addons
func (collection Collection) Count() int {
	return len(collection.addons)
}

// Add Add an Addon to the collection
func (collection *Collection) Add(addon Addon) {
	collection.addons = append(collection.addons, addon)
}

// NewCollection create a new addon instance
func NewCollection() Collection {
	var addons Collection

	return addons
}
