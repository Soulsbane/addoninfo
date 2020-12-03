package main

import (
	"fmt"

	"github.com/alexflint/go-arg"
)

// ListCommand used for various type of list addons to the CLI
type ListCommand struct {
	Command     string `arg:"positional"`
	SetUpstream bool   `arg:"-u"`
}

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
}
