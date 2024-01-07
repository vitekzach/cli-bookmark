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
	Categories []string
	Commands   []Command
	Shells     []Shell
}

type Config interface {
	save()
	backup()
	//TODO: verity integrity
}

func readconfigvalues() configValues {
	Logger.Debug("Loading config values")

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
		Logger.Warn(fmt.Sprintf("Unknown system %v", runtime.GOOS))
		panic(fmt.Sprintf("You are using an unsupported operating system: %v. %v", runtime.GOOS, defaultErrorAppendage))
		// TODO how to handle this in the future?
	}

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		Logger.Info("Config file does not exist, saving a default one")
		conf.save()
		Logger.Debug(fmt.Sprintf("Loaded config: %+v", conf))
		return conf
	}

	Logger.Debug("Config file exists, loading it")

	configJson, err := os.ReadFile(configFilePath)

	if err != nil {
		Logger.Error("Config couldn't be read from a file", "error", err)
		panic(fmt.Sprintf("Your config file located in %v cannot be read, fix or remove it. %v", configFilePath, defaultErrorAppendage))
		// TODO replace with a warning and do not panic
	}

	err = json.Unmarshal(configJson, &conf)
	if err != nil {
		Logger.Error("Config couldn't unmarshaled from JSON", "error", err)
		panic(fmt.Sprintf("Your config file located in %v cannot be parsed as JSON, fix or remove it. %v", configFilePath, defaultErrorAppendage))
		// TODO replace with a warning and do not panic
	}

	Logger.Debug(fmt.Sprintf("Loaded config: %+v", conf))
	conf.backup()

	return conf
}

func (c *configValues) saveconfigas(filePath string) {
	Logger.Debug(fmt.Sprintf("Saving config values to: %v", filePath))
	configJson, err := json.Marshal(c)

	if err != nil {
		Logger.Error("Config couldn't be converted to json", "error", err)
		fmt.Printf("WARNING: Error while saving config file (cannot marshal JSON), your settings will not be persisted. %v", defaultErrorAppendage)
		// TODO make it red?
	}

	err = os.WriteFile(filePath, configJson, 0666)

	if err != nil {
		Logger.Error(fmt.Sprintf("Config couldn't be saved to config path %v", configFilePath), "error", err)
		fmt.Printf("WARNING: Cannot save config file, your settings will not be persisted. %v", defaultErrorAppendage)
		// TODO error here
		// TODO make it red
	}

	Logger.Debug("Config saved")
}

func (c *configValues) save() {
	Logger.Debug("Saving config regular way")
	c.saveconfigas(configFilePath)
}

func (c *configValues) backup() {
	Logger.Info("Backing up config")
	c.saveconfigas(configBackupFilePath)
}

func establishFolderPaths() {
	// switch runtime.GOOS {
	// case "windows":
	// 	configFolder = `%APPDATA%\CliBookmark`
	// default:
	// 	//TODO
	// 	configFolder = "TODO"
	// }

	userConfigFolder, err := os.UserConfigDir()
	if err != nil {
		Logger.Error("User config dir could not be established!", "error", err)
		// TODO raise my own error? Panic here?
	}

	var configLeafFolder string

	switch runtime.GOOS {
	case "windows":
		configLeafFolder = "CliBookmark"
	default:
		//TODO
		configLeafFolder = "clibookmark"
	}

	configFolder = filepath.Join(userConfigFolder, configLeafFolder)
	configFilePath = filepath.Join(configFolder, "config.json")
	configBackupFilePath = filepath.Join(configFolder, "config_backup.json")

	defaultErrorAppendage = fmt.Sprintf("Logs and config can be found in %v. Please raise and issue or contribute at: %v", configFolder, repoLink)

	//TODO move GetConfig code to here

}

func GetConfig() {
	initConsts()
	establishFolderPaths()

	Logger.Debug(fmt.Sprintf("Config folder location: %v", configFolder))

	if _, err := os.Stat(configFolder); os.IsNotExist(err) {
		Logger.Info("Config folder does not exist, creating...")

		err = os.MkdirAll(configFolder, 0666)
		if err != nil {
			Logger.Error("Failed to make config folder", "error", err)
			// TODO panic here?
		}

		Logger.Info("Config folder created")
	}

	config := readconfigvalues()

	// config.updateconfigvalues(config)

	Logger.Info(fmt.Sprintf("Loaded config for app version %v", config.Version))
	Logger.Info(fmt.Sprintf("Config app version %v", config.Version))
	// config.saveconfigvalues()
}
