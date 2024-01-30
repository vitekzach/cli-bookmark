package src

import (
	"os"
	"runtime"
	"testing"
)

func TestInitConsts(t *testing.T) {
	expectedDefaultCategories := []Category{{Name: "Default"}}
	exptectedCategoriesStr := "[{\"Name\": \"Not Default\"}]"
	exptectedCategories := []Category{{Name: "Not Default"}}
	categoriesWrongInput := "[bla]"

	os.Setenv("CLI_BOOKMARK_CATEGORIES", "")
	err := initConsts()
	if err != nil {
		t.Fatalf("initConsts threw following error: %v", err.Error())
	}
	t.Logf("%+v, %+v", defaultCategories, expectedDefaultCategories)
	if defaultCategories[0] != expectedDefaultCategories[0] {
		t.Fatalf("Expected default category 0 %+v, got %+v", expectedDefaultCategories[0], defaultCategories[0])
	}
	if len(defaultCategories) != len(expectedDefaultCategories) {
		t.Fatalf("Exptected lentgth of categories %v, got %v", len(defaultCategories), len(expectedDefaultCategories))
	}
	if len(defaultCommands) != 0 {
		t.Fatalf("Expected length of default commands 0, got %v", len(defaultCommands))
	}

	os.Setenv("CLI_BOOKMARK_CATEGORIES", exptectedCategoriesStr)
	err = initConsts()
	if err != nil {
		t.Fatalf("initConsts threw following error: %v", err.Error())
	}
	if defaultCategories[0] != exptectedCategories[0] {
		t.Fatalf("Expected category 0 %+v, got %+v", expectedDefaultCategories[0], expectedDefaultCategories[0])
	}
	if len(defaultCategories) != len(exptectedCategories) {
		t.Fatalf("Exptected lentgth of categories %v, got %v", len(defaultCategories), len(exptectedCategories))
	}

	os.Setenv("CLI_BOOKMARK_CATEGORIES", categoriesWrongInput)
	err = initConsts()
	if err == nil {
		t.Fatalf("initConsts did not throw an error, exptected one")
	}
	if defaultCategories[0] != expectedDefaultCategories[0] {
		t.Fatalf("Expected default category 0 %+v, got %+v", expectedDefaultCategories[0], defaultCategories[0])
	}
	if len(defaultCategories) != len(expectedDefaultCategories) {
		t.Fatalf("Exptected lentgth of categories %v, got %v", len(defaultCategories), len(expectedDefaultCategories))
	}

	if defaultCategories[0].Name != "Default" {
		t.Fatalf("Expected 1st category to be %v, got %v", "Default", defaultCategories[0].Name)
	}

	switch runtime.GOOS {
	case "windows":
		{

			if len(defaultShells) != 2 {
				t.Fatalf("Length of Windows shells should be 0, instead got %v.", len(defaultShells))
			}
		}
	case "linux":
		{
			if len(defaultShells) != 0 {
				t.Fatalf("Length of Linux shells should be 0, instead got %v.", len(defaultShells))
			}
		}
	case "ios":
		{
			if len(defaultShells) != 0 {
				t.Fatalf("Length of iOS shells should be 0, instead got %v.", len(defaultShells))
			}
		}

	}

	// check when passing defaults by env vars
	os.Setenv("CLI_BOOKMARK_CATEGORIES", "[{\"Name\": \"Default\"}, {\"Name\":\"Custom\"}]")

	err = initConsts()
	if err != nil {
		t.Fatalf("initConsts threw following error: %v", err.Error())
	}
	if len(defaultCategories) != 2 {
		t.Fatalf("Expected len 2 of categories, got %v", len(defaultCategories))
	}

	if defaultCategories[0].Name != "Default" {
		t.Fatalf("Expected 1st category to be %v, got %v", "Default", defaultCategories[0].Name)
	}

	if defaultCategories[1].Name != "Custom" {
		t.Fatalf("Expected 1st category to be %v, got %v", "Custom", defaultCategories[0].Name)
	}

	// add default shells

	if len(defaultCommands) != 0 {
		t.Fatalf("Expected length of defaultCommands 0, got %v", len(defaultCommands))
	}

	os.Setenv("CLI_BOOKMARK_CATEGORIES", "XXXX")
	err = initConsts()
	if err == nil {
		t.Fatal("Categories shouldn't be able to be unmarshalled")
	}
}
