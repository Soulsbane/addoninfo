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
}
