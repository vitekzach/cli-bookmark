package src

var currentVersion string
var repoLink string
var defaultCategories []string
var defaultLinuxShells []Shell
var defaultIosShells []Shell
var defaultWindowsShells []Shell
var defaultCommands []Command

func initConsts() {
	currentVersion = "0.0.1"
	repoLink = "https://github.com/vitekzach/cli-bookmark"

	defaultCategories = []string{"Default"}
	defaultLinuxShells = []Shell{}
	defaultIosShells = []Shell{}
	defaultWindowsShells = []Shell{}
	defaultCommands = []Command{}

}