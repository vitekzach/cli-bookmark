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

type configValues struct {
	Version string
}

type Config interface {
	saveconfigvalues()
}

func readconfigvalues() configValues {
	// TODO what if config doesn't exist yet?
	Logger.Debug("Loading config values.")

	// create default config
	var conf configValues
	conf.Version = "0.0.1"

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		Logger.Info("Config file does not exist, saving a default one.")
		conf.saveconfigvalues()
		return conf
		// TODO save config
	}
	// TODO make a config backup first

	Logger.Debug("Config file exists, loading it")

	configJson, err := os.ReadFile(configFilePath)

	if err != nil {
		Logger.Error("Config couldn't be read from a file", "error", err)
		panic(fmt.Sprintf("Your config file located in %v cannot be read, fix or remove it.", configFilePath))
	}

	err = json.Unmarshal(configJson, &conf)
	if err != nil {
		Logger.Error("Config couldn't unmarshaled from JSON", "error", err)
		panic(fmt.Sprintf("Your config file located in %v cannot be parsed as JSON, fix or remove it.", configFilePath))
	}

	Logger.Debug("Config loaded")

	return conf
}

func (c *configValues) saveconfigvalues() {
	Logger.Debug("Saving config values.")
	configJson, err := json.Marshal(c)

	if err != nil {
		Logger.Error("Config couldn't be converted to json", "error", err)
		fmt.Printf("WARNING: Error while saving config file (cannot marshal JSON), your settings will not be persisted.")
		// TODO make it red?
		//TODO reference log
		//TODO reference github
	}

	err = os.WriteFile(configFilePath, configJson, 0666)

	if err != nil {
		Logger.Error(fmt.Sprintf("Config couldn't be saved to config path %v", configFilePath), "error", err)
		fmt.Printf("WARNING: Cannot save config file, your settings will not be persisted.")
		// TODO error here
		// TODO make it red
		//TODO reference log
		//TODO reference github
	}

	Logger.Debug("Config saved")

}

func init() {
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

}

func GetConfig() {
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
	config.saveconfigvalues()
}
