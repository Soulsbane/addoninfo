package main

import (
	"addoninfo/addons"
	"fmt"

	"github.com/alexflint/go-arg"
)

func main() {
	var args struct {
		List *ListCommand `arg:"subcommand:list"`
		All  bool         `arg:"-q"`
	}

	arg.MustParse(&args)

	switch {
	case args.List != nil:
		fmt.Printf("List Command:  %s\n", args.List.Command)
	}

	addon := addons.NewAddon()
	addon.TestParser()
	fmt.Println("Addons Author: ", addon.GetAuthor())

	collection := addons.NewCollection()
	collection.Add(addon)
	fmt.Println(collection.Count())
	parser := addons.NewTocParser()
	data := "## Author: Paul\n## Description: Does it work\n"
	parser.ParseString(data)
	parser.Dump()
}
