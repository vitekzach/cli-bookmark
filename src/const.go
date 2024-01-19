package src

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

	defaultCategories = []Category{Category{Name: "Default"}}
	defaultLinuxShells = []Shell{}
	defaultIosShells = []Shell{}
	defaultWindowsShells = []Shell{
		{Name: "Command prompt", ShortName: "cmd", Command: "cmd.exe"}, // TODO
		{Name: "PowerShell", ShortName: "ps1", Command: "ps"},          // TODO
	}
	defaultCommands = []Command{}

}
