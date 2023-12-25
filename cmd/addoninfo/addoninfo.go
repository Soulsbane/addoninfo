package main

import (
	"addoninfo/addons"
	"addoninfo/internal/config"
	"errors"
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/erikgeiser/promptkit/confirmation"
	"github.com/ncruces/zenity"
)

var configFile = config.New()

func promptToAddPath() {
	input := confirmation.New("No paths configured. Add a new path now?", confirmation.Yes)
	yesNo, err := input.RunPrompt()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		if yesNo {
			handleAddPathCommand()
		} else {
			fmt.Println("A new path can be added with the 'add-path' command.")
		}
	}
}

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
