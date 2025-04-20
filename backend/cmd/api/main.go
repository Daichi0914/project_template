package main

import (
	"database/sql"
	"fmt"
	"github.com/Daichi0914/project_template/infrastructure/db"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

func retryConnect(connStr string, retries int, delay time.Duration) (*sql.DB, error) {
	for i := 0; i < retries; i++ {
		projectDb, err := sql.Open("postgres", connStr)
		if err == nil {
			if err := projectDb.Ping(); err == nil {
				return projectDb, nil
			}
		}
		log.Printf("Waiting for DB... (%d/%d)", i+1, retries)
		time.Sleep(delay)
	}
	return nil, fmt.Errorf("failed to connect to DB after %d attempts", retries)
}

func main() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
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
	defer func(dbConn *sql.DB) {
		err := dbConn.Close()
		if err != nil {

		}
	}(dbConn)

	db.RunMigration(dbConn)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, "ok")
		if err != nil {
			return
		}
	})

	fmt.Println("Listening on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
