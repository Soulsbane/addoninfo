package main

// ListCommand used for various type of list addons to the CLI
type ListCommand struct {
	Command string `arg:"positional"`
}
