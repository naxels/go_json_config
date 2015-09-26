package config

import (
	"os"
	"reflect"
	"testing"
)

const (
	testDataDir    = "testdata"
	testFileSuffix = ".json"
)

type Config struct {
	URL      string
	User     string
	Projects []string
}

func TestParse(t *testing.T) {
	var c Config

	fileName := "config1"
	fileLocation := testDataDir + string(os.PathSeparator) + fileName + testFileSuffix

	expectedURL := "https://github.com"
	expectedUser := "naxels"
	expectedProjects := []string{"go_csv_file_checker", "markdown_http_server"}

	err := Parse(fileLocation, &c)
	if err != nil {
		t.Fatalf("Error opening/parsing file %v", fileName)
	}

	if c.URL != expectedURL {
		t.Fatalf("Expected URL: %v, got URL: %v", expectedURL, c.URL)
	}

	if c.User != expectedUser {
		t.Fatalf("Expected Projects: %v, got Projects: %v", expectedUser, c.User)
	}

	if reflect.DeepEqual(c.Projects, expectedProjects) == false {
		t.Fatalf("Expected Projects: %v, got Projects: %v", expectedProjects, c.Projects)
	}
}
