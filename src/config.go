package src

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

var configFolder string

type configValues struct {
	version string
}

type Config interface {
	//TODO
	getconfigvalues() configValues
	updateconfigvalues(configValues)
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

	config := configValues{version: "0.0.1"}
	Logger.Info(fmt.Sprintf("Loaded config for app version %v", config.version))
}
