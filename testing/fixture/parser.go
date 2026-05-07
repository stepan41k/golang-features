package main

import (
	"fmt"
	"os"
	"testing"
)

func TetstParse(t *testing.T) {
	data, err := os.ReadFile("testdata/input.json")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(data)atomic
	
}