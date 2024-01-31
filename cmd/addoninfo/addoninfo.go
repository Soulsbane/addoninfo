package main

import (
	"addoninfo/addons"
	"addoninfo/internal/config"
	"errors"
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/ncruces/zenity"
)

var configFile = config.New()

func handleAddPathCommand() {
	path, err := zenity.SelectFile(zenity.Filename(""), zenity.Directory())

	if err != nil {
		if errors.Is(err, zenity.ErrCanceled) {
			fmt.Println("User Canceled. No path added.")
		} else {
			fmt.Println("Error: ", err)
		}
	} else {
		configFile.AddInstallPath(path)
		configFile.Save()

		if err != nil {
			fmt.Println("Error saving config file: ", err)
		} else {
			fmt.Println("Added path: ", path)
		}
	}
}

func main() {
	var cmd commands

	arg.MustParse(&cmd)
	collection := addons.NewCollection()
	collection.Build(".")

	switch {
	case cmd.List != nil:
		collection.List(cmd.List.Type)
	case cmd.AddPath != nil:
		handleAddPathCommand()
	}

}
