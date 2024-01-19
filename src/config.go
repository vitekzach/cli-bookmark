package src

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

var configFolder string
var configFilePath string
var configBackupFilePath string
var defaultErrorAppendage string

type configValues struct {
	Version    string
	Categories []Category
	Commands   []Command
	Shells     []Shell
}

type Config interface {
	save()
	backup()
	//TODO: verity integrity
	verifyIntegrity()
	// TODO: add command
	AddCommand()
	// TODO: add category
	AddCategory()
	// TODO: remove command
	RemoveCommand()
	// TODO: remove category
	RemoveCategory()
}

func readconfigvalues() configValues {
	logger.Debug("Loading config values")

	// create default config
	var conf configValues
	conf.Version = currentVersion
	conf.Categories = defaultCategories
	conf.Commands = defaultCommands

	switch runtime.GOOS {
	case "windows":
		conf.Shells = defaultWindowsShells
	case "linux":
		conf.Shells = defaultLinuxShells
	case "ios":
		conf.Shells = defaultIosShells

	default: // unkown system
		logger.Warn(fmt.Sprintf("Unknown system %v", runtime.GOOS))
		error(fmt.Sprintf("You are using an unsupported operating system: %v. %v", runtime.GOOS, defaultErrorAppendage))
		os.Exit(1)
	}

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		logger.Info("Config file does not exist, saving a default one")
		conf.save()
		logger.Debug(fmt.Sprintf("Loaded config: %+v", conf))
		return conf
	}

	logger.Debug("Config file exists, loading it")

	configJson, err := os.ReadFile(configFilePath)

	if err != nil {
		logger.Error("Config couldn't be read from a file", "error", err)
		error(fmt.Sprintf("Your config file located in %v cannot be read, fix or remove it. %v", configFilePath, defaultErrorAppendage))
		os.Exit(1)
	}

	err = json.Unmarshal(configJson, &conf)
	if err != nil {
		logger.Error("Config couldn't unmarshaled from JSON", "error", err)
		err = os.WriteFile(configBackupFilePath, configJson, 0666)
		if err != nil {
			logger.Error("Config backup couldn't be saved.", "error", err)
		}
		error(fmt.Sprintf("Your config file located in %v cannot be parsed as JSON, fix or remove it. There should be a backup already created for you in %v. %v", configFilePath, configBackupFilePath, defaultErrorAppendage))
		os.Exit(1)
	}

	logger.Debug(fmt.Sprintf("Loaded config: %+v", conf))
	conf.backup()

	return conf
}

func (c *configValues) saveconfigas(filePath string) {
	logger.Debug(fmt.Sprintf("Saving config values to: %v", filePath))
	configJson, err := json.Marshal(c)

	if err != nil {
		logger.Error("Config couldn't be converted to json", "error", err)
		warning(fmt.Sprintf("Error while saving config file (cannot marshal JSON), your settings will not be persisted! %v", defaultErrorAppendage))
		return
	}

	err = os.WriteFile(filePath, configJson, 0666)

	if err != nil {
		logger.Error(fmt.Sprintf("Config couldn't be saved to config path %v", configFilePath), "error", err)
		warning(fmt.Sprintf("Cannot save config file, your settings will not be persisted! %v", defaultErrorAppendage))
		return
	}

	logger.Debug("Config saved")
}

func (c *configValues) save() {
	logger.Debug("Saving config regular way")
	c.saveconfigas(configFilePath)
}

func (c *configValues) backup() {
	logger.Info("Backing up config")
	c.saveconfigas(configBackupFilePath)
}

func establishFolderPaths() {
	userConfigFolder, err := os.UserConfigDir()
	if err != nil {
		logger.Error("User config dir could not be established!", "error", err)
		error(fmt.Sprintf("Your user config folder failed to be established. %v", defaultErrorAppendage))
		os.Exit(1)
	}

	var configLeafFolder string

	switch runtime.GOOS {
	case "windows":
		configLeafFolder = "CliBookmark"
	default:
		configLeafFolder = "clibookmark"
	}

	configFolder = filepath.Join(userConfigFolder, configLeafFolder)
	configFilePath = filepath.Join(configFolder, "config.json")
	configBackupFilePath = filepath.Join(configFolder, "config_backup.json")

	defaultErrorAppendage = fmt.Sprintf("Logs and config can be found in %v. Please raise an issue or contribute at: %v", configFolder, repoLink)
}

func getConfig() {
	// TODO once proper interface for bookmarker is made, do not export this
	initConsts()
	establishFolderPaths()

	logger.Debug(fmt.Sprintf("Config folder location: %v", configFolder))

	if _, err := os.Stat(configFolder); os.IsNotExist(err) {
		logger.Info("Config folder does not exist, creating...")

		err = os.MkdirAll(configFolder, 0666)
		if err != nil {
			logger.Error("Failed to make config folder", "error", err)
			error(fmt.Sprintf("Your config folder couldn't be created. %v", defaultErrorAppendage))
			os.Exit(1)
		}

		logger.Info("Config folder created")
	}

	config := readconfigvalues()

	logger.Info(fmt.Sprintf("Loaded config for app version %v", config.Version))
	logger.Info(fmt.Sprintf("Config app version %v", config.Version))
}
