package config

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig("test_config.yaml")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if config.IODocument.Header.DocumentTitle == "" {
		t.Error("Expected DocumentTitle to be set, but it was empty")
	}
}
