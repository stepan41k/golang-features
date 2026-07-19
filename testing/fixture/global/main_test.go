package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Starting Docker with Postgres...")

	code := m.Run()

	fmt.Println("Stoping Docker...")

	os.Exit(code)
}