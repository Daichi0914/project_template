package bootstrap

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func retryConnect(connStr string, retries int, delay time.Duration) (*sql.DB, error) {
	for i := 0; i < retries; i++ {
		db, err := sql.Open("postgres", connStr)
		if err == nil && db.Ping() == nil {
			return db, nil
		}
		log.Printf("Waiting for DB... (%d/%d)", i+1, retries)
		time.Sleep(delay)
	}
	return nil, fmt.Errorf("failed to connect to DB after %d attempts", retries)
}

func InitDB() *sql.DB {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	dbConn, err := retryConnect(connStr, 10, 3*time.Second)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}
	return dbConn
}
