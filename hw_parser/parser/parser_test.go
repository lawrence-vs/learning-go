package parser

import (
	"testing"
)

func TestParseExcel(t *testing.T) {
	data, err := ParseExcel("test_input.xlsx")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(data) == 0 {
		t.Error("Expected parsed data, but got empty slice")
	}
}
