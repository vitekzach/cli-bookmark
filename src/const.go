package src

import (
	"errors"
	"fmt"
	"os"
	"runtime"
)

var currentVersion string
var repoLink string
var defaultCategories []Category
var defaultShells []Shell
var defaultCommands []Command

func initConsts() error {
	currentVersion = "0.0.1"
	repoLink = "https://github.com/vitekzach/cli-bookmark"

	defaultCategoriesFallback := []Category{{Name: "Default"}}
	defaultCategories = defaultCategoriesFallback
	_, err := setFromEnvJSONWithDefaultStruct[[]Category]("CLI_BOOKMARK_CATEGORIES", defaultCategoriesFallback, &defaultCategories)
	if err != nil {
		return errors.New("value in environmetal varible CLI_BOOKMARK_CATEGORIES cannot be unmashalled")
	}

	switch runtime.GOOS {
	case "windows":
		defaultShells = []Shell{
			{Name: "Command prompt", Command: "cmd.exe"}, // TODO
			{Name: "PowerShell", Command: "ps"},          // TODO
		}
	case "linux":
		defaultShells = defaultShells
	case "ios":
		defaultShells = defaultShells
	default: // unkown system
		logger.Warn(fmt.Sprintf("Unknown system %v", runtime.GOOS))
		printError(fmt.Sprintf("You are using an unsupported operating system: %v. %v", runtime.GOOS, defaultErrorAppendage))
		os.Exit(1)
	}

	defaultCommands = []Command{}

	return nil

}
