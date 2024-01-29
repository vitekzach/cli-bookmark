package src

import (
	"encoding/json"
	"fmt"
	"os"
)

func setFromEnvWithDefaultStr(envVarName string, defaultVal string, valToSet *string) bool {
	envVarValue := os.Getenv(envVarName)
	if envVarValue != "" {
		*valToSet = envVarValue
		return true
	} else {
		*valToSet = defaultVal
		return false
	}
}

func setFromEnvJSONWithDefaultStruct[T any](envVarName string, defaultVal T, valToSet *T) (bool, error) {
	var envVarContent string
	varEnvPresent := setFromEnvWithDefaultStr(envVarName, "", &envVarContent)
	if !varEnvPresent {
		*valToSet = defaultVal
		return false, nil
	}

	err := json.Unmarshal([]byte(envVarContent), &valToSet)
	if err != nil {
		return true, fmt.Errorf("value in environmetal varible %v cannot be unmashalled", envVarName)
	}
	return true, nil

}
