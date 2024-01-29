package src

import (
	"os"
	"testing"
)

func TestSetFromEnvWithDefault(t *testing.T) {
	defaultVal := "default.json"
	var varToSet string

	setFromEnvWithDefaultStr("CLI_BOOKMARK_DEFAULT", defaultVal, &varToSet)

	if varToSet != defaultVal {
		t.Fatalf("Expected %v, got \"%v\"", defaultVal, varToSet)
	}

	envToSet := "not_default.json"

	os.Setenv("CLI_BOOKMARK_DEFAULT", envToSet)
	setFromEnvWithDefaultStr("CLI_BOOKMARK_DEFAULT", defaultVal, &varToSet)

	if varToSet != envToSet {
		t.Fatalf("Expected %v, got \"%v\"", envToSet, varToSet)
	}
}

func TestSetFromEnvJSONWithDefaultStruct(t *testing.T) {
	type testStruct struct {
		Val int
	}

	testEnvVar := "CLI_BOOKMARK_DEFAULT"
	os.Setenv(testEnvVar, "")

	var varToSet testStruct
	defaultVal := testStruct{9}
	valPresent, error := setFromEnvJSONWithDefaultStruct[testStruct](testEnvVar, defaultVal, &varToSet)

	if valPresent {
		t.Fatalf("Expected env var not to be present, but it is.")
	}
	if error != nil {
		t.Fatalf(error.Error())
	}

	if varToSet != defaultVal {
		t.Fatalf("Expected %v, got \"%v\"", defaultVal, varToSet)
	}

	envToSet := "r"
	os.Setenv("CLI_BOOKMARK_DEFAULT", envToSet)
	valPresent, error = setFromEnvJSONWithDefaultStruct[testStruct](testEnvVar, defaultVal, &varToSet)
	if error == nil {
		t.Fatalf("Expected unmashall error, but it didn't error.")
	}

	envToSet = "{\"Val\": 5}"

	os.Setenv("CLI_BOOKMARK_DEFAULT", envToSet)
	valPresent, error = setFromEnvJSONWithDefaultStruct[testStruct]("CLI_BOOKMARK_DEFAULT", defaultVal, &varToSet)

	if !valPresent {
		t.Fatalf("Expected env var to be present, but it is not.")
	}
	if error != nil {
		t.Fatalf(error.Error())
	}
	expectedVar := testStruct{5}
	if varToSet != expectedVar {
		t.Fatalf("Expected %v, got \"%v\"", expectedVar, varToSet)
	}
}
