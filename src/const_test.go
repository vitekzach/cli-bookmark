package src

import (
	"os"
	"testing"
)

func TestInitConsts(t *testing.T) {
	// check when run by default
	initConsts()

	if defaultCategories[0].Name != "Default" {
		t.Fatalf("Expected 1st category to be %v, got %v", "Default", defaultCategories[0].Name)
	}

	// check when passing defaults by env vars
	os.Setenv("CLI_BOOKMARK_CATEGORIES", "Default,Custom")
	initConsts()

	if len(defaultCategories) != 2 {
		t.Fatalf("Expected len 2 of categories, got %v", len(defaultCategories))
	}

	if defaultCategories[0].Name != "Default" {
		t.Fatalf("Expected 1st category to be %v, got %v", "Default", defaultCategories[0].Name)
	}

	if defaultCategories[1].Name != "Custom" {
		t.Fatalf("Expected 1st category to be %v, got %v", "Custom", defaultCategories[0].Name)
	}

	// TODO check linux shells
	// TODO check windows shells
	// TODO check IOS shells
	// TODO check default commands
}
