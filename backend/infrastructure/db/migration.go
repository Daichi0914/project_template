package db

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
)

func RunMigration(db *sql.DB) {
	sqlPath := filepath.Join("infrastructure", "db", "init.sql")
	content, err := os.ReadFile(sqlPath)
	if err != nil {
		log.Fatalf("Failed to read migration file: %v", err)
	}

	_, err = db.Exec(string(content))
	if err != nil {
		log.Fatalf("Failed to execute migration: %v", err)
	}
}
