package main

import (
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

	parser := NewTocParser()
	parser.AddEntry("Version", "1.0")
	parser.AddEntry("Author", "Soulsbane")
	parser.AddEntry("Name", "TocParser")
	parser.Dump()
	fmt.Println("HasEntry for Version: ", parser.HasEntry("Version"))
	fmt.Println("HasEntry for VersionSSSSSSSSS: ", parser.HasEntry("VersionSSSSSSSSS"))
	fmt.Println(parser.GetEntry("Author"))
	fmt.Println(parser.GetEntry("AuthorZZZZ"), "<<<")
}
