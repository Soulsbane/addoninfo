package main

import (
	"addoninfo/addons"
	"github.com/alexflint/go-arg"
)

func main() {
	var cmd commands

	arg.MustParse(&cmd)
	collection := addons.NewCollection()
	collection.Build(".")

	switch {
	case cmd.List != nil:
		collection.List(cmd.List.Type)
	}

}
