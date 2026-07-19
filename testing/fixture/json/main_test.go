package main

import (
	"os"
	"testing"
)

func loadFixture(t *testing.T, filename string) []byte {
	data, err := os.ReadFile("testdata/" + filename)
	
	if err != nil {
		t.Fatalf("Failed to load fixture: %v", err)
	}

	return data
} 