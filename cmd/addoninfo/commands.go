package main

// ListCommand used for various type of list addons to the CLI
type ListCommand struct {
	Type string `arg:"positional"`
}

type AddPathCommand struct {
}

type commands struct {
	AddPath *AddPathCommand `arg:"subcommand:add-path" help:"Add a path to a WoW install to the list "`
	List    *ListCommand    `arg:"subcommand:list" help:"Show a list of unused saved variable or list of installed addons"`
}
