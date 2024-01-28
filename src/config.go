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

// type error interface {
//     Error() string
// }

type Config interface {
	save()
	backup()
	//TODO: verity integrity
	verifyIntegrity()
	// TODO: add command
	addCommand() error
	// TODO: add category
	addCategory()
	// TODO: remove command
	removeCommand()
	// TODO: remove category
	removeCategory()
}

func (c *configValues) addCommand(comm Command) error {
	c.Commands = append(c.Commands, comm)
	return nil
}

func readconfigvalues() configValues {
	logger.Debug("Loading config values")

	// create default config
	var conf configValues
	conf.Version = currentVersion
	conf.Categories = defaultCategories
	conf.Commands = defaultCommands

	conf.Shells = defaultShells

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
		printError(fmt.Sprintf("Your config file located in %v cannot be read, fix or remove it. %v", configFilePath, defaultErrorAppendage))
		os.Exit(1)
	}

	err = json.Unmarshal(configJson, &conf)
	if err != nil {
		logger.Error("Config couldn't unmarshaled from JSON", "error", err)
		err = os.WriteFile(configBackupFilePath, configJson, 0666)
		if err != nil {
			logger.Error("Config backup couldn't be saved.", "error", err)
		}
		printError(fmt.Sprintf("Your config file located in %v cannot be parsed as JSON, fix or remove it. There should be a backup already created for you in %v. %v", configFilePath, configBackupFilePath, defaultErrorAppendage))
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
		printWarning(fmt.Sprintf("Error while saving config file (cannot marshal JSON), your settings will not be persisted! %v", defaultErrorAppendage))
		return
	}

	err = os.WriteFile(filePath, configJson, 0666)

	if err != nil {
		logger.Error(fmt.Sprintf("Config couldn't be saved to config path %v", configFilePath), "error", err)
		printWarning(fmt.Sprintf("Cannot save config file, your settings will not be persisted! %v", defaultErrorAppendage))
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
		printError(fmt.Sprintf("Your user config folder failed to be established. %v", defaultErrorAppendage))
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

	var configFilenamePath string
	configFileNameEnv := os.Getenv("CLI_BOOKMARK_CONFIG_FILENAME")
	if configFileNameEnv != "" {
		configFilenamePath = configFileNameEnv
	} else {
		configFilenamePath = "config.json"
	}

	var configBackupFilenamePath string
	configBackupFileNameEnv := os.Getenv("CLI_BOOKMARK_CONFIG_BACKUP_FILENAME")
	if configBackupFileNameEnv != "" {
		configBackupFilenamePath = configBackupFileNameEnv
	} else {
		configBackupFilenamePath = "config_backup.json"
	}

	configFilePath = filepath.Join(configFolder, configFilenamePath)
	configBackupFilePath = filepath.Join(configFolder, configBackupFilenamePath)

	defaultErrorAppendage = fmt.Sprintf("Logs and config can be found in %v. Please raise an issue or contribute at: %v", configFolder, repoLink)
}

func getConfig() {
	// TODO once proper interface for bookmarker is made, do not export this
	err := initConsts()
	if err != nil {
		panic(err.Error())
	}
	establishFolderPaths()

	logger.Debug(fmt.Sprintf("Config folder location: %v", configFolder))

	if _, err := os.Stat(configFolder); os.IsNotExist(err) {
		logger.Info("Config folder does not exist, creating...")

		err = os.MkdirAll(configFolder, 0666)
		if err != nil {
			logger.Error("Failed to make config folder", "error", err)
			printError(fmt.Sprintf("Your config folder couldn't be created. %v", defaultErrorAppendage))
			os.Exit(1)
		}

		logger.Info("Config folder created")
	}

	config := readconfigvalues()

	logger.Info(fmt.Sprintf("Loaded config for app version %v", config.Version))
	config.addCommand(Command{Command: "dir", Category: config.Categories[0]})
	logger.Debug(fmt.Sprintf("After add: %+v", config))

	logger.Info(fmt.Sprintf("Config app version %v", config.Version))
}
