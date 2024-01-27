package src

import (
	"os"
	"strings"
)

var currentVersion string
var repoLink string
var defaultCategories []Category
var defaultLinuxShells []Shell
var defaultIosShells []Shell
var defaultWindowsShells []Shell
var defaultCommands []Command

func initConsts() {
	currentVersion = "0.0.1"
	repoLink = "https://github.com/vitekzach/cli-bookmark"

	defaultEnvCategories := os.Getenv("CLI_BOOKMARK_CATEGORIES")
	if defaultEnvCategories != "" {
		defaultCategories = []Category{}
		for _, cat := range strings.Split(defaultEnvCategories, ",") {
			defaultCategories = append(defaultCategories, Category{Name: cat})
		}
	} else {

		defaultCategories = []Category{{Name: "Default"}}
	}

	defaultLinuxShells = []Shell{}
	defaultIosShells = []Shell{}
	defaultWindowsShells = []Shell{
		{Name: "Command prompt", ShortName: "cmd", Command: "cmd.exe"}, // TODO
		{Name: "PowerShell", ShortName: "ps1", Command: "ps"},          // TODO
	}
	defaultCommands = []Command{}

}
