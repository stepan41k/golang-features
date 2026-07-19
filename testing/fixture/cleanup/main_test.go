package main

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, _ := sql.Open("postgres", "test_dsn")

	t.Cleanup(func() {
		db.Close()
		fmt.Println("Database connection closed")
	})

	return db
}

func TestMyService(t *testing.T) {
	db := setupTestDB(t)

	assert.NotNil(t, db)
}