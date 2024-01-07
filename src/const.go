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
	// TODO replace with real version
	repoLink = "https://github.com/something"

	defaultCategories = []string{"Default"}
	defaultLinuxShells = []Shell{}
	defaultIosShells = []Shell{}
	defaultWindowsShells = []Shell{}
	defaultCommands = []Command{}

}
