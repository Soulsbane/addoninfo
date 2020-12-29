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
	collection := addons.NewCollection()
	collection.Build(".")

	switch {
	case args.List != nil:
		fmt.Printf("List Command:  %s\n", args.List.Command)
		collection.List(args.List.Command)
	}

}
