package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	processor, err := NewProcessor("input.txt")
	if err != nil {
		t.Fatal("Error creating code processor: ", err)
	}
	defer processor.File.Close()

	expected := 281
	result := processor.ProcessCodes()
	if result != expected {
		t.Errorf("Expected log output to contain %d, got %d", expected, result)
	}
}
