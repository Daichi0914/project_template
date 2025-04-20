package main

import (
	"database/sql"
	"fmt"
	"github.com/Daichi0914/project_template/backend/db"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	connStr := "host=db port=5432 user=admin password=admin dbname=project_template sslmode=disable"
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
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
