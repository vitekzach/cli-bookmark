package src

import (
	"fmt"
	"testing"
)

func TestAddCommand(t *testing.T) {
	testCategory := Category{Name: "Default"}
	commandToAdd := Command{Command: "dir", Category: testCategory}
	testConfigValues := configValues{Version: "0.01", Categories: []Category{testCategory}, Commands: []Command{}, Shells: defaultIosShells}

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
