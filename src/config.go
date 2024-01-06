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
	readconfigvalues() configValues
	updateconfigvalues(configValues)
}

func (c configValues) readconfigvalues() configValues {
	Logger.Debug("Loading config values.")

	configJson, err := os.ReadFile(configFilePath)

	if err != nil {
		Logger.Error("Config couldn't be read from a file")
		// TODO error here
	}

	var config configValues

	err = json.Unmarshal(configJson, &config)
	if err != nil {
		Logger.Error("Config couldn't unmarshaled from JSON")
		// TODO error here
	}

	Logger.Debug("Config loaded.")

	return config
}

func (c configValues) updateconfigvalues(conf configValues) {
	Logger.Debug("Saving config values.")
	configJson, err := json.Marshal(conf)

	if err != nil {
		Logger.Error("Config couldn't be converted to json")
		// TODO error here
	}

	err = os.WriteFile(configFilePath, configJson, 0666)

	if err != nil {
		Logger.Error(fmt.Sprintf("Config couldn't be saved to config path %v", configFilePath))
		// TODO error here
	}

	Logger.Debug("Config saved.")

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
			Logger.Error(err.Error())
			// TODO panic here?
		}

		Logger.Info("Config folder created.")
	}

	config := configValues{}
	config = config.readconfigvalues()
	config.updateconfigvalues(config)

	Logger.Info(fmt.Sprintf("Loaded config for app version %v", config.Version))
}
