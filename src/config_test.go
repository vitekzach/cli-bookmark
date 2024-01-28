package src

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestAddCommand(t *testing.T) {
	testCategory := Category{Name: "Default"}
	commandToAdd := Command{Command: "dir", Category: testCategory}
	testConfigValues := configValues{Version: "0.01", Categories: []Category{testCategory}, Commands: []Command{}, Shells: defaultShells}

	for i := 1; i < 5; i++ {
		testname := fmt.Sprintf("%+v", commandToAdd)
		t.Run(testname, func(t *testing.T) {
			testConfigValues.addCommand(commandToAdd)
			if i != len(testConfigValues.Commands) {
				t.Errorf("Length of commands %v, expected %v", i, len(testConfigValues.Commands))
			}
		})
	}
}

func TestEstablishFolderPaths(t *testing.T) {
	establishFolderPaths()
	userConfigDir, _ := os.UserConfigDir()
	var configLeafFolder string
	switch runtime.GOOS {
	case "windows":
		configLeafFolder = "CliBookmark"
	default:
		configLeafFolder = "clibookmark"
	}
	expectedConfigPath := filepath.Join(userConfigDir, configLeafFolder, "config.json")
	if configFilePath != expectedConfigPath {
		t.Errorf("configFilePath %v does not match expected path %v", configFilePath, expectedConfigPath)
	}

	extectedBackupPath := filepath.Join(userConfigDir, configLeafFolder, "config_backup.json")
	if configBackupFilePath != extectedBackupPath {
		t.Errorf("configBackupFilePath %v does not match expected path %v", configBackupFilePath, extectedBackupPath)
	}

	os.Setenv("CLI_BOOKMARK_CONFIG_FILENAME", "test_config.json")
	os.Setenv("CLI_BOOKMARK_CONFIG_BACKUP_FILENAME", "test_config_backup.json")
	establishFolderPaths()
	expectedConfigPath = filepath.Join(userConfigDir, configLeafFolder, "test_config.json")
	if configFilePath != expectedConfigPath {
		t.Errorf("configFilePath %v does not match expected path %v", configFilePath, expectedConfigPath)
	}

	extectedBackupPath = filepath.Join(userConfigDir, configLeafFolder, "test_config_backup.json")
	if configBackupFilePath != extectedBackupPath {
		t.Errorf("configBackupFilePath %v does not match expected path %v", configBackupFilePath, extectedBackupPath)
	}
}
